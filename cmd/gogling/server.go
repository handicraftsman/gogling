package main

import (
  "context"
  "net/http"
  "log"
  "path"
  "time"
  "plugin"
  "sync"
  "fmt"

  "github.com/gorilla/mux"
  lua "github.com/yuin/gopher-lua"
  luar "layeh.com/gopher-luar"
)

// Gogling - stores information about current gogling instance
type Gogling struct {
  AI *AppInfo
  Lctx *lua.LState
  Logger *log.Logger
  Router *mux.Router
  Plugins *sync.Map
  LoadedPlugins *sync.Map
  Server *http.Server
  SC chan bool
}

// GoglingI - an instance of Gogling
var GoglingI *Gogling

func Serve(ai *AppInfo) {
  l := log.New(StdoutFile, "server ", -1)
  l.Printf("Starting")

  GoglingI = &Gogling{
    AI: ai,
    Lctx: nil,
    Logger: l,
    Router: nil,
    Plugins: &sync.Map{},
    LoadedPlugins: &sync.Map{},
    Server: nil,
    SC: make(chan bool, 16),
  }
  defer func() {
    if GoglingI.Lctx != nil {
      GoglingI.Lctx.Close()
    }
  }()
  go LoadLua(nil)
  GoglingI.SC <- true
}

func LoadLua(oldLua *lua.LState) {
  <-GoglingI.SC

  if oldLua != nil {
    oldLua.Close()
  }
  GoglingI.Lctx = lua.NewState()

  if (GoglingI.Server != nil) {
    ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(15 * time.Second))
    defer cancel()
    GoglingI.Server.Shutdown(ctx)
  }
  GoglingI.Router = mux.NewRouter()

  GoglingI.Plugins = &sync.Map{}
  GoglingI.LoadedPlugins = &sync.Map{}
  g := GoglingI.Lctx.NewTable()
  GoglingI.Lctx.SetGlobal("_GoglingI", luar.New(GoglingI.Lctx, GoglingI))
  GoglingI.Lctx.SetGlobal("gogling", g)
  GoglingI.Lctx.SetGlobal("_reload", luar.New(GoglingI.Lctx, func(s map[string]interface{}, msg string) {
    fmt.Fprintln(s["writer"].(http.ResponseWriter), msg)
    s["to_reload"] = GoglingI.Lctx
  }))
  GoglingI.Lctx.SetGlobal("_wrapSession", luar.New(GoglingI.Lctx, func(f func(map[string]interface{})) func(writer http.ResponseWriter, request *http.Request) {
    return func(writer http.ResponseWriter, request *http.Request) {
      s := map[string]interface{}{
        "writer": writer,
        "request": request,
      }
      defer func() {
        if r := recover(); r != nil {
          http.Error(writer, "500 Internal Server Error", http.StatusInternalServerError)
        }
      }()
      f(s)
      if s["to_reload"] != nil {
        writer.(http.Flusher).Flush()
      }
      if s["to_reload"] != nil {
        go LoadLua(s["to_reload"].(*lua.LState))
        GoglingI.SC <- true
      }
    }
  }))
  GoglingI.Lctx.SetGlobal("import_go", luar.New(GoglingI.Lctx, func(n string) {
    _, ok := GoglingI.LoadedPlugins.Load(n)
    if !ok {
      _, ok = GoglingI.Plugins.Load(n)
      if !ok {
        p, err := plugin.Open(path.Join(GoglingI.AI.Config.PluginDir, "go2lua_" + n + ".so"))
        if err != nil {
          GoglingI.Logger.Panic(err)
        }
        i, err := p.Lookup("GoglingLoad")
        if err != nil {
          GoglingI.Logger.Panic(err)
        }
        i.(func(L *lua.LState))(GoglingI.Lctx)
        GoglingI.Plugins.Store(n, i)
      }
    }
  }))
  if err := GoglingI.Lctx.DoString(luaUtils); err != nil {
    GoglingI.Logger.Panic(err)
  }

  if err := GoglingI.Lctx.DoFile(path.Join(GoglingI.AI.Config.RootDir, GoglingI.AI.Config.RouterFile)); err != nil {
    log.Panic(err)
  }

  GoglingI.Server = &http.Server{
    Handler: GoglingI.Router,
    Addr: GoglingI.AI.Config.Host+":"+GoglingI.AI.Config.Port,
    WriteTimeout: 15 * time.Second,
    ReadTimeout:  15 * time.Second,
  }
  GoglingI.Server.ListenAndServe()
}

const luaUtils string = `
  gogling.I = _GoglingI
  gogling.U = {}
  gogling.U.reload = _reload
  gogling.U.wrap = _wrapSession
  gogling.U.import = function(n)
    import_go(n)
    return require('go.' .. n)
  end
`