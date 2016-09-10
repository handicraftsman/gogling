/* main.go
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

import "flag"

var sName = "Gogling"
var sVer = "0.0.1-pre3"
var sDone = make(chan bool)
var sTestName *string
var sAllTests bool

// Main Function, sir!
func main() {
	sTestName = flag.String("test", "none", "Test selector") // To allow running tests
	flag.Parse()                                             // Parse flags
	//tRunTests()

	lInit()
	lMain.Printf("It's %s v%s", sName, sVer) // Output info about our Gogling
	lMain.Println("Started")

	go cMain() // Load config in separate thread
	<-sDone    // Wait until config loads

	sDone = make(chan bool) // To not exit before net-thread finishes
	go nMain()              // Start network thread

	<-sDone // Send true here to exit
	lMain.Println("# Main: Exiting")
}
