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
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yuin/gopher-lua"
)

func pProcess(iWrt http.ResponseWriter, iReq *http.Request, iData string, iPath string) /*(string, int)*/ {
	lFile := fGetInfo(iPath) // Get info about file
	lType, lTypeG := pGetType(lFile.Ext)

	if lFile.Ext == ".lua" { // If target file is lua script, execute it
		if pLuaParse(lFile.Name) {
			hlErr(iWrt, iReq, iPath, 500) // If lua errored - crash
		}
	} else if lTypeG == "text" {
		lData, err := ioutil.ReadFile("data/" + lFile.Name) // Read file
		if !hlErrScan(iWrt, iReq, lFile.Name, err) {        // If not errored
			iWrt.Header().Set("Content-Type", "text/"+lType+"; charset=utf-8")
			iWrt.Header().Set("X-Content-Type-Options", "nosniff")
			iWrt.Header().Set("Server", sName+" "+sVer)

			iWrt.WriteHeader(200) // Set code

			fmt.Fprint(iWrt, string(lData)) // Send data
		}
	} else if lTypeG == "raw" {
		lData, err := ioutil.ReadFile("data/" + lFile.Name)
		if !hlErrScan(iWrt, iReq, lFile.Name, err) {
			iWrt.Header().Set("Content-Type", lType+"; charset=utf-8")
			iWrt.Header().Set("X-Content-Type-Options", "nosniff")
			iWrt.Header().Set("Server", sName+" "+sVer)

			iWrt.WriteHeader(200) // Set code

			lBuf := new(bytes.Buffer)
			err2 := binary.Write(lBuf, binary.LittleEndian, lData)
			checkRuntimeErr(lPrep, err2)

			iWrt.Write(lBuf.Bytes()) // Send data
		}
	}
}

func pGetType(iExt string) (string, string) {
	var iType string
	var iTypeG string

	var lText = "text"
	var lRaw = "raw"

	switch iExt {
	case ".html":
		iType = "html"
		iTypeG = lText
		break
	case ".txt":
		iType = lText
		iTypeG = lText
		break
	case ".md":
		iType = "markdown"
		iTypeG = lText
		break

	case ".png":
		iType = "image/png"
		iTypeG = lRaw
		break
	case ".jpg":
		iType = "image/jpeg"
		iTypeG = lRaw
		break
	case ".gif":
		iType = "image/gif"
		iTypeG = lRaw
		break

	default:
		iType = "application/octet-stream"
		iTypeG = lRaw
		break
	}

	return iType, iTypeG
}

func pLuaParse(iPath string) bool { // Lua runner
	lLua := lua.NewState()              // Init Lua
	defer lLua.Close()                  // Close VM after finishing
	mMain(lLua)                         // Load modules
	err := lLua.DoFile("data/" + iPath) // Run needed file

	return checkRuntimeErr(lPrep, err) // Check for errors
}
