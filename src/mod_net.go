package main

import (
	"fmt"
	"io/ioutil"

	lua "github.com/yuin/gopher-lua"
)

func mNetLoader(lLua *lua.LState) int {
	mod := lLua.SetFuncs(lLua.NewTable(), mNetExports) // Register Funtions
	lLua.Push(mod)                                     // Return Module

	return 1
}

var mNetExports = map[string]lua.LGFunction{ // Here we are storing functions
	"init":           mNetInit,
	"echo":           mNetSend,
	"send_file_html": mNetSendFileHTML,
	"close":          mNetClose,
}

func mNetInit(iLua *lua.LState) int { // Init HTTP message
	gWriter.Header().Set("Content-Type", "text/html; charset=utf-8")
	gWriter.Header().Set("X-Content-Type-Options", "nosniff")

	return 0
}

func mNetSend(iLua *lua.LState) int { // Send data
	iData := iLua.ToString(1) // Get argument
	fmt.Fprint(gWriter, iData)
	return 0
}

func mNetSendFileHTML(iLua *lua.LState) int { // File Sender
	iPath := iLua.ToString(1)                      // Get argument
	lData, err := ioutil.ReadFile("data/" + iPath) // Read file
	if checkRuntimeErr(lLuaP, err) {               // If errored - crash
		iLua.RaiseError("Cannot read file: %s\n", "data/"+iPath)
	} else { // Else - send file
		fmt.Fprint(gWriter, string(lData))
	}
	return 0
}

func mNetClose(iLua *lua.LState) int { // Exit
	gWriter.WriteHeader(200)
	return 0
}
