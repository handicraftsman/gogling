/* handlers.go
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
	"os"
	"strconv"
)

// Now this (and next) variable can be accessed from everywhere
var gWriter http.ResponseWriter
var gRequest *http.Request

// Sender
func hSend(iWrt http.ResponseWriter, iData string, iCode int) {
	// Set content type to HTML
	iWrt.Header().Set("Content-Type", "text/html; charset=utf-8")
	iWrt.Header().Set("X-Content-Type-Options", "nosniff")

	iWrt.WriteHeader(iCode) // Send response code

	fmt.Fprintf(iWrt, iData) // Send response
}

// Errorer
func hlErr(iWrt http.ResponseWriter, iReq *http.Request, iGet string, iCode int) {
	lPath := "err/" + strconv.Itoa(iCode) + ".html" // Get correct path

	lData, err := ioutil.ReadFile(lPath) // Get error file (if any)
	var lOut string
	if !os.IsNotExist(err) { // Check error file for existance
		lOut = string(lData)
	} else { // If there's no such file - send default text
		lOut = "Gogling says: \"" + strconv.Itoa(iCode) + "\""
	}

	hSend(iWrt, lOut, iCode)                                 // Send data
	lNet.Printf("\033[31m# %d: data/%s\033[0m", iCode, iGet) // Send notification into console
}

// Error-Scanner
func hlErrScan(iWrt http.ResponseWriter, iReq *http.Request, iGet string, iErr error) bool {
	if os.IsNotExist(iErr) { // Handle 404 (not found)
		hlErr(iWrt, iReq, iGet, 404)
		return true
	} else if os.IsPermission(iErr) { // Handle 403 (access denied)
		hlErr(iWrt, iReq, iGet, 403)
		return true
	}
	return false
}

// Gogling-info handler
func hGoglingInfo(iWrt http.ResponseWriter, iReq *http.Request) {
	if sConf["webInfoEnabled"] == "true" { // Is this allowed?
		// Send info
		fmt.Fprintf(iWrt, "<style>body{padding-left:16px}p{padding-left:32px}</style>")
		fmt.Fprintf(iWrt, "<h1>Gogling info:</h1>\n")
		fmt.Fprintf(iWrt, "<p>")
		fmt.Fprintf(iWrt, "Version: %s\n", sVer)
		fmt.Fprintf(iWrt, "</p>")
	} else { // HEY!
		lNet.Println("\033[31m# Somebody tried to access info page, but it's disabled in config\033[0m")
		fmt.Fprintf(iWrt, "<h1><strong>Sorry, this output is disabled in gogling's config</strong></h1>")
	}
}

// Main handler. Gets files
func hMain(iWrt http.ResponseWriter, iReq *http.Request) {
	fmt.Println("ListenAttemp")
	gWriter = iWrt // To make it accessible
	gRequest = iReq

	nURL = iReq.URL
	nURLData = nURL.Query()

	var lGet = iReq.URL.Path[1:]   // To make life easier
	if lGet == "" || lGet == "/" { // To allow "/" requests
		lGet = sConf["index"]
	}

	lData, err := ioutil.ReadFile("data/" + lGet) // Get data & error (if any)
	lErrored := hlErrScan(iWrt, iReq, lGet, err)  // Scan for errors

	if !lErrored { // If no errors - send data
		pProcess(iWrt, iReq, string(lData), lGet) // Send data
		//lNet.Println("\033[32m# Sent:", lGet, "\033[0m") // Send notification into console
	}

	gWriter = nil // Clear me!
	gRequest = nil
	nURL = nil
	nURLData = nil
}

// Network thread
func nMain() {
	http.HandleFunc("/", hMain)                                       // Add file access to it
	http.HandleFunc("/hGoglingInfo", hGoglingInfo)                    // Gogling info. Will show if enabled in config
	lNet.Println("Listening on " + sConf["ip"] + ":" + sConf["port"]) // Tell about it
	http.ListenAndServe(sConf["ip"]+":"+sConf["port"], nil)           // Listen
}
