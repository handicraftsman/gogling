package main

import (
  "net/http"
  "log"
  "path"
  "time"

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
  }
  defer GoglingI.Lctx.Close()
  g := GoglingI.Lctx.NewTable()
  GoglingI.Lctx.SetField(g, "I", luar.New(GoglingI.Lctx, GoglingI))
  GoglingI.Lctx.SetGlobal("gogling", g)
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
  gogling.U = {}
  gogling.U.wrap = function(f)
    return function(writer, request)
      local session = { writer = writer, request = request }
      f(session)
    end
  end
`