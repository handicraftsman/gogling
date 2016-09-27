/* mod_net.go
 *
 * Copyright (C) 2016 Nickolay Ilyushin <nickolay02@inbox.ru>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"bytes"
	"encoding/binary"
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
	"init":       mNetInit,
	"echo":       mNetSend,
	"sendf_text": mNetSendText,
	"sendf_raw":  mNetSendRaw,
	"close":      mNetClose,
}

func mNetInit(iLua *lua.LState) int { // Init HTTP message
	iType := iLua.ToString(1)
	gWriter.Header().Set("Content-Type", iType+"; charset=utf-8")
	gWriter.Header().Set("X-Content-Type-Options", "nosniff")
	gWriter.Header().Set("Server", sName+" "+sVer)
	return 0
}

func mNetSend(iLua *lua.LState) int { // Send data
	iData := iLua.ToString(1) // Get argument
	fmt.Fprint(gWriter, iData)
	return 0
}

func mNetSendText(iLua *lua.LState) int { // File Sender
	iPath := iLua.ToString(1)                      // Get argument
	lData, err := ioutil.ReadFile("data/" + iPath) // Read file
	if checkRuntimeErr(lLuaP, err) {               // If errored - crash
		iLua.RaiseError("Cannot read file: %s\n", "data/"+iPath)
	} else { // Else - send file
		fmt.Fprint(gWriter, string(lData))
	}
	return 0
}

func mNetSendRaw(iLua *lua.LState) int {
	iPath := iLua.ToString(1)
	lData, err := ioutil.ReadFile("data/" + iPath)
	if checkRuntimeErr(lLuaP, err) {
		iLua.RaiseError("Cannot read file: %s\n", "data/"+iPath)
	} else {
		lBuf := new(bytes.Buffer)
		err2 := binary.Write(lBuf, binary.LittleEndian, lData)
		checkWarn(lPrep, err2)

		gWriter.Write(lBuf.Bytes()) // Send data
	}
	return 0
}

func mNetGetPost(iLua *lua.LState) int {
	iName := iLua.ToString(1)
	oValue := gRequest.PostFormValue(iName)
	iLua.Push(lua.LString(oValue))
	return 0
}

func mNetClose(iLua *lua.LState) int { // Exit
	iCode := iLua.ToInt(1)
	gWriter.WriteHeader(iCode)
	return 0
}
