/* mod_misc.go
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
	luar "github.com/layeh/gopher-luar"
	lua "github.com/yuin/gopher-lua"
)

func mNetRestLoader(iLua *lua.LState) int {
	mod := iLua.SetFuncs(iLua.NewTable(), mNetRestExports) // Register Functions
	iLua.Push(mod)                                         // Return Module
	return 1
}

var mNetRestExports = map[string]lua.LGFunction{ // Here we are storing our funcs
	"method": mNetRestMethod,

	"get":    mNetRestGet,
	"getAll": mNetRestGetAll,

	"post":    mNetRestPost,
	"postAll": mNetRestPostAll,
}

/* UTILS */
func mNetRestMethod(iLua *lua.LState) int {
	lMethod := gRequest.Method      // Get method
	iLua.Push(lua.LString(lMethod)) // Return it
	return 1
}

/* GET */
func mNetRestGet(iLua *lua.LState) int {
	lKey := iLua.ToString(1)     // Get key
	lValue := nURLData.Get(lKey) // Get value using key
	if lValue == "" {            // Set to "nil" if it's empty
		lValue = "#nil"
	}
	iLua.Push(lua.LString(lValue)) // Return value
	return 1
}

func mNetRestGetAll(iLua *lua.LState) int {
	lOut := make(map[string]string) // Make for-output map
	for lKey := range nURLData {    // For each key in our GET:
		lValue := nURLData.Get(lKey) // Get value using key
		lOut[lKey] = lValue          // Save it to output map
	}
	iLua.Push(luar.New(iLua, lOut)) // Return result
	return 1
}

/* POST */
func mNetRestPost(iLua *lua.LState) int {
	lKey := iLua.ToString(1)               // Get key
	lValue := gRequest.PostFormValue(lKey) // Get value using key
	if lValue == "" {                      // Set to "nil" if it's empty
		lValue = "#nil"
	}
	iLua.Push(lua.LString(lValue)) // Return value
	return 1
}

func mNetRestPostAll(iLua *lua.LState) int {
	lOut := make(map[string]string)       // Make for-output map
	for lKey := range gRequest.PostForm { // For each key in our POST:
		lValue := gRequest.PostForm.Get(lKey) // Get value using key
		lOut[lKey] = lValue                   // Save it to output map
	}
	iLua.Push(luar.New(iLua, lOut)) // Return result
	return 1
}
