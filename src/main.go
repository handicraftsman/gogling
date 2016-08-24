package main

import "log"

var sName = "gogling"
var sVer = "0.0.0"
var sWebInfoEnabled = true
var sDone = make(chan bool)

// Main Function, sir!
func main() {
	log.Print("# Main: Started")

	go nMain() // Start network thread

	<-sDone
	log.Print("# Main: Exiting")
}
