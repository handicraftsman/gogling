/* pre_main.go
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
	"io/ioutil"
	"net/http"

	"github.com/yuin/gopher-lua"
)

func pProcess(iWrt http.ResponseWriter, iReq *http.Request, iData string, iPath string) /*(string, int)*/ {
	lFile := fGetInfo(iPath) // Get info about file
	if lFile.Ext == ".lua" { // If target file is lua script, execute it
		if pLuaParse(lFile.Name) {
			hlErr(iWrt, iReq, iPath, 500) // If lua errored - crash
		}
	} else {
		lData, err := ioutil.ReadFile("data/" + lFile.Name) // Read file
		if !hlErrScan(iWrt, iReq, lFile.Name, err) {        // If not errored
			iWrt.Header().Set("Content-Type", "text/html; charset=utf-8")
			iWrt.Header().Set("X-Content-Type-Options", "nosniff")

			iWrt.WriteHeader(200) // Set code

			fmt.Fprint(iWrt, string(lData)) // Send data
		}
	}
}

func pLuaParse(iPath string) bool { // Lua runner
	lLua := lua.NewState()              // Init Lua
	defer lLua.Close()                  // Close VM after finishing
	mMain(lLua)                         // Load modules
	err := lLua.DoFile("data/" + iPath) // Run needed file

	return checkRuntimeErr(lPrep, err) // Check for errors
}
