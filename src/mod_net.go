package main

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

func mNetLoader(lLua *lua.LState) int {
	// register functions to the table
	mod := lLua.SetFuncs(lLua.NewTable(), mNetExports)
	// register other stuff
	lLua.SetField(mod, "name", lua.LString("Boo"))

	// returns the module
	lLua.Push(mod)

	return 1
}

var mNetExports = map[string]lua.LGFunction{
	"init":  mNetInit,
	"echo":  mNetSend,
	"close": mNetClose,
}

func mNetInit(iLua *lua.LState) int {
	gWriter.Header().Set("Content-Type", "text/html; charset=utf-8")
	gWriter.Header().Set("X-Content-Type-Options", "nosniff")

	return 0
}

func mNetSend(iLua *lua.LState) int {
	iData := iLua.ToString(1)
	fmt.Fprint(gWriter, iData)
	return 0
}

func mNetClose(iLua *lua.LState) int {
	gWriter.WriteHeader(200)
	return 0
}
