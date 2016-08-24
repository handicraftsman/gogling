/* handlers.go
 *
 * Copyright (C) 2015-2016 Nickolay Ilyushin <nickolay02@inbox.ru>
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
	"log"
	"net/http"
	"os"
	"strconv"
)

// Error-sender
func hlErrSend(iWrt http.ResponseWriter, iData string, iCode int) {
	// Set content type to HTML
	iWrt.Header().Set("Content-Type", "text/html; charset=utf-8")
	iWrt.Header().Set("X-Content-Type-Options", "nosniff")

	// Send response code
	iWrt.WriteHeader(iCode)
	// Send response
	fmt.Fprintf(iWrt, iData)
}

// Error-Handler
func hlErr(iWrt http.ResponseWriter, iReq *http.Request, iGet string, iCode int) {
	lPath := "err/" + strconv.Itoa(iCode) + ".html"            // Get correct path
	log.Println("# Net: Errored. Sending", lPath, "to client") // Notify

	lData, err := ioutil.ReadFile(lPath) // Get error file (if any)
	var lOut string
	if !os.IsNotExist(err) { // Check error file for existance
		lOut = string(lData)
	} else { // If there's no such file - send default text
		lOut = "Gogling says: \"404 page not found\""
	}

	hlErrSend(iWrt, lOut, http.StatusNotFound) // Send data
	log.Printf("# Net: %d: %s", iCode, iGet)   // Send notification into console
}

// Error-Scanner
func hlErrScan(iWrt http.ResponseWriter, iReq *http.Request, iGet string, iErr error) bool {
	if os.IsNotExist(iErr) { // Handle 404
		hlErr(iWrt, iReq, iGet, 404)
		return true
	} //else if os.isPermission(err) {
	//}
	return false
}

// Gogling-info handler
func hGoglingInfo(iWrt http.ResponseWriter, iReq *http.Request) {
	if sWebInfoEnabled { // Is this allowed?
		// Send info
		fmt.Fprintf(iWrt, "<style>body{padding-left:16px}p{padding-left:32px}</style>")
		fmt.Fprintf(iWrt, "<h1>Gogling info:</h1>\n")
		fmt.Fprintf(iWrt, "<p>")
		fmt.Fprintf(iWrt, "Version: %s\n", sVer)
		fmt.Fprintf(iWrt, "</p>")
	} else { // HEY!
		log.Println("# Net: Somebody tried to access info page, but it's disabled in config")
		fmt.Fprintf(iWrt, "<h1><strong>Sorry, this output is disabled in gogling's config</strong></h1>")
	}
}

// Main handler. Gets files
func hMain(iWrt http.ResponseWriter, iReq *http.Request) {
	var lGet = iReq.URL.Path[1:]   // To make life easier
	if lGet == "" || lGet == "/" { // To allow "/" requests
		lGet = "index.html"
	}

	lData, err := ioutil.ReadFile("data/" + lGet) // Get data & error (if any)
	lErrored := hlErrScan(iWrt, iReq, lGet, err)  // Scan for errors

	if !lErrored { // If no errors - send data
		// Set content type to HTML
		iWrt.Header().Set("Content-Type", "text/html; charset=utf-8")
		iWrt.Header().Set("X-Content-Type-Options", "nosniff")

		fmt.Fprintf(iWrt, string(lData))  // Send data
		log.Println("# Net: Sent:", lGet) // Send notification into console
	}
}

// Network thread
func nMain() {
	http.HandleFunc("/", hMain)
	http.HandleFunc("/hGoglingInfo", hGoglingInfo)
	http.ListenAndServe(":8080", nil)
}
