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

import "path"

type fInfo struct {
	Name string
	Ext  string
}

func fGetInfo(iName string) fInfo {
	var oOut fInfo
	var err error

	oOut.Name = iName
	oOut.Ext = path.Ext(iName)
	checkRuntimeErr(lFile, err)

	return oOut
}

/**
func fCheckMatch(iName string, iRegExp string) bool {
	oMatch, err := regexp.MatchString(iRegExp, iName) // Check text
	checkWarn(lRegex, err)                            // Check for errors
	return oMatch                                     // Return result
}

func fRunCmd(iBin string, iFile string) string {
	lCmd := exec.Command(iBin, iFile) // Declare process
	lCmd.Dir = "data/"                // Change working directory
	oData, err := lCmd.Output()       // Run, wait for finishing and get output
	checkWarn(lFile, err)             // Check for errors
	return string(oData)              // Return output
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
	checkWarn(lFile, err)                             // Check for errors
	return oOut                                       // Return output
}
**/
