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
