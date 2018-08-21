package main

import (
  "net/http"
  "log"
  "path"
  "time"
  "plugin"
  "sync"

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
}

// GoglingI - an instance of Gogling
var GoglingI *Gogling

func Serve(ai *AppInfo) {
  l := log.New(StdoutFile, "server ", -1)
  l.Printf("Starting")

  GoglingI = &Gogling{
    AI: ai,
    Lctx: lua.NewState(),
    Logger: l,
    Router: mux.NewRouter(),
    Plugins: &sync.Map{},
  }
  defer GoglingI.Lctx.Close()
  g := GoglingI.Lctx.NewTable()
  GoglingI.Lctx.SetGlobal("_GoglingI", luar.New(GoglingI.Lctx, GoglingI))
  GoglingI.Lctx.SetGlobal("gogling", g)
  GoglingI.Lctx.SetGlobal("import_go", luar.New(GoglingI.Lctx, func(n string) {
    _, ok := GoglingI.Plugins.Load(n)
    if !ok {
      p, err := plugin.Open(path.Join(GoglingI.AI.Config.PluginDir, "go2lua_" + n + ".so"))
      if err != nil {
        GoglingI.Logger.Fatal(err)
      }
      i, err := p.Lookup("GoglingLoad")
      if err != nil {
        GoglingI.Logger.Fatal(err)
      }
      i.(func(L *lua.LState))(GoglingI.Lctx)
      GoglingI.Plugins.Store(n, i)
    }
  }));
  if err := GoglingI.Lctx.DoString(luaUtils); err != nil {
    GoglingI.Logger.Fatal(err)
  }

  if err := GoglingI.Lctx.DoFile(path.Join(ai.Config.RootDir, ai.Config.RouterFile)); err != nil {
    log.Fatal(err)
  }

  s := &http.Server{
    Handler: GoglingI.Router,
    Addr: ai.Config.Host+":"+ai.Config.Port,
    WriteTimeout: 15 * time.Second,
    ReadTimeout:  15 * time.Second,
  }
  go s.ListenAndServe()
}

const luaUtils string = `
  gogling.I = _GoglingI
  gogling.U = {}
  gogling.U.wrap = function(f)
    return function(writer, request)
      local session = { writer = writer, request = request }
      f(session)
    end
  end
  gogling.U.import = function(n)
    import_go(n)
    return require('go.' .. n)
  end
`