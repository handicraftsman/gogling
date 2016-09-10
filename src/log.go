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
