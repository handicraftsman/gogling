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
	"html/template"
	"net/http"
)

// Returns output to iWrt, input: iData, iPName
func pProcess(iWrt http.ResponseWriter, iData string, iPName string) {
	lFile := fGetInfo(iPName)

	if lFile.Type == "html" {
		lTmpl, err := template.New(iPName).Parse(iData) // Parse input
		errC := checkParseErr("Preprocessor", err)
		if errC {
			hlErr(iWrt, nil, iPName, 500)
		}

		err = lTmpl.Execute(iWrt, template.HTML("")) // Execute template
		errC = checkParseErr("Preprocessor", err)
		if errC {
			hlErr(iWrt, nil, iPName, 500)
		}
	} else if lFile.IsTemplate {
		fmt.Fprintf(iWrt, fRunCmd(lFile.Type, lFile.Name))
	}

	// Done!
}

/* Why so short?
 *
 * Go has it's own preprocessor in 'text/template' and 'html/template' packages
 * We are using them here
 */
