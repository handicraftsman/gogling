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
	"fmt"

	luar "github.com/layeh/gopher-luar"
	lua "github.com/yuin/gopher-lua"
)

func mNetRestLoader(iLua *lua.LState) int {
	mod := iLua.SetFuncs(iLua.NewTable(), mNetRestExports) // Register Functions
	iLua.Push(mod)                                         // Return Module
	return 1
}

var mNetRestExports = map[string]lua.LGFunction{ // Here we are storing our funcs
	"get":    mNetRestGet,
	"getAll": mNetRestGetAll,
}

// GET
func mNetRestGet(iLua *lua.LState) int {
	lKey := iLua.ToString(1)

	lValue := nURLData.Get(lKey)
	if lValue == "" {
		lValue = "nil"
	}

	iLua.Push(lua.LString(lValue))
	return 1
}

func mNetRestGetAll(iLua *lua.LState) int {
	var lOut = make(map[string]string)
	fmt.Println("Start-Parse")
	for lKey := range nURLData {
		lValue := nURLData.Get(lKey)
		fmt.Println(lKey + " | " + lValue)
		lOut[lKey] = lValue
	}
	fmt.Println("End-Parse")
	iLua.Push(luar.New(iLua, lOut))
	return 1
}
