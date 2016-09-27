/* mod.go
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
