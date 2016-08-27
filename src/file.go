/* file.go
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
	"os/exec"
	"regexp"
)

type fInfo struct {
	Name       string
	Type       string
	IsTemplate bool
}

func fGetInfo(iName string) fInfo {
	var oOut fInfo
	oOut.Name = iName
	oOut.Type = fGetFileType(iName)
	oOut.IsTemplate = fIsTemplate(iName)
	return oOut
}

func fCheckMatch(iName string, iRegExp string) bool {
	oMatch, err := regexp.MatchString(iRegExp, iName) // Check text
	checkWarn("RegExp", err)                          // Check for errors
	return oMatch                                     // Return result
}

func fRunCmd(iBin string, iFile string) string {
	lCmd := exec.Command(iBin, "data/"+iFile) // Run process
	oData, err := lCmd.Output()               // Wait for finishing and get output
	checkWarn("File Runner", err)             // Check for errors
	return string(oData)                      // Return output
}

func fGetFileType(iName string) string {
	var oOut string                   // Define output variable
	if fCheckMatch(iName, "\\.lua") { // For Lua
		oOut = "lua"
	} else { // If can't detect
		oOut = "plain-text"
	}
	return oOut // Return output
}

func fIsTemplate(iName string) bool {
	var oOut bool                                     // Define output variable
	oOut, err := regexp.MatchString("\\.tmpl", iName) // Ckeck text
	checkWarn("RegExp", err)                          // Check for errors
	return oOut                                       // Return output
}
