package main

import lua "github.com/yuin/gopher-lua"

func mMiscLoader(iLua *lua.LState) int {
	mod := iLua.SetFuncs(iLua.NewTable(), mMiscExports) // Register Functions
	iLua.Push(mod)                                      // Return Module

	return 1
}

var mMiscExports = map[string]lua.LGFunction{ // Here we are storing our funcs
	"log": mMiscLog,
}

func mMiscLog(iLua *lua.LState) int { // Sends string from 1st arg to stdout
	iText := iLua.ToString(1) // Get first argument
	lLuaP.Print(iText)        // Print it
	return 0
}
