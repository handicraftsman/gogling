/* log.go
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
	"io"
	"log"
	"os"
	"syscall"
)

var lStdout io.Writer // Declare stdout

var lMain *log.Logger // Declare logs
var lNet *log.Logger
var lPrep *log.Logger
var lRegex *log.Logger
var lConf *log.Logger
var lFile *log.Logger
var lLog *log.Logger
var lLuaP *log.Logger

func lInit() { // Init all wars
	lStdout = os.NewFile(uintptr(syscall.Stdout), "stdout")

	lMain = log.New(lStdout, "# Main: ", -1)
	lNet = log.New(lStdout, "# Net: ", -1)
	lPrep = log.New(lStdout, "# Preprocessor: ", -1)
	lRegex = log.New(lStdout, "# RegExp: ", -1)
	lConf = log.New(lStdout, "# Config: ", -1)
	lFile = log.New(lStdout, "# File: ", -1)
	lLuaP = log.New(lStdout, "# Lua: ", -1)
}

func checkErr(iLog *log.Logger, iErr error) { // Error checker
	if iErr != nil {
		iLog.Printf("\033[31m%s \033[0m\n", iErr.Error())
		os.Exit(1)
	}
}

func checkWarn(iLog *log.Logger, iErr error) { // Warning checker
	if iErr != nil {
		iLog.Printf("\033[33m%s \033[0m\n", iErr.Error())
	}
}

func checkRuntimeErr(iLog *log.Logger, iErr error) bool { // Runtime-error checker
	if iErr != nil {
		iLog.Printf("\033[33m%s \033[0m\n", iErr.Error())
		return true
	}
	return false
}

func printToLog(iLog *log.Logger, iText string) { // Not really needed.
	if iLog != nil {
		iLog.Println(iText)
	} else {
		log.Panicln("# Log: logger is nil!")
	}
}
