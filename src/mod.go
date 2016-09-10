package main

import lua "github.com/yuin/gopher-lua"

func mMain(iLua *lua.LState) {
	iLua.PreloadModule("gogling", mGLoader)         // Include main gogling's module
	iLua.PreloadModule("gogling.net", mNetLoader)   // Network module
	iLua.PreloadModule("gogling.misc", mMiscLoader) // Misc module
}

func mGLoader(iLua *lua.LState) int {
	mod := iLua.SetFuncs(iLua.NewTable(), mGExports) // Register Functions

	iLua.SetField(mod, "name", lua.LString(sName))   // Stores gogling's name :P
	iLua.SetField(mod, "version", lua.LString(sVer)) // Same with version

	iLua.Push(mod) // Return Module

	return 1
}

var mGExports = map[string]lua.LGFunction{} // Empty function map
