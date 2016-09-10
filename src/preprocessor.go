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
			hlErr(iWrt, iReq, iPath, 500)
		}
	} else {
		lData, err := ioutil.ReadFile("data/" + lFile.Name)
		if !hlErrScan(iWrt, iReq, lFile.Name, err) {
			iWrt.Header().Set("Content-Type", "text/html; charset=utf-8")
			iWrt.Header().Set("X-Content-Type-Options", "nosniff")

			iWrt.WriteHeader(200)

			fmt.Fprint(iWrt, string(lData))
		}

	}

}

func pLuaParse(iPath string) bool {
	lLua := lua.NewState()
	defer lLua.Close()
	mMain(lLua)
	err := lLua.DoFile("data/" + iPath)

	return checkRuntimeErr(lPrep, err)
}

/* OLD
import (
	"fmt"
	"html/template"
	"net/http"
)

// Returns output to iWrt, input: iData, iPName
func pProcess(iWrt http.ResponseWriter, iData string, iPName string) {
	lFile := fGetInfo(iPName)

	if lFile.Type == "html" {
		lTmpl, err := template.New(iPName).Parse(iData) // Parse input
		errC := checkRuntimeErr(lPrep, err)
		if errC {
			hlErr(iWrt, nil, iPName, 500)
		}

		err = lTmpl.Execute(iWrt, template.HTML("")) // Execute template
		errC = checkRuntimeErr(lPrep, err)
		if errC {
			hlErr(iWrt, nil, iPName, 500)
		}
	} else if lFile.IsTemplate {
		fmt.Fprintf(iWrt, fRunCmd(lFile.Type, lFile.Name))
	}

	// Done!
}

/**/
/* Why so short?
 *
 * Go has it's own preprocessor in 'text/template' and 'html/template' packages
 * We are using them here
 */
