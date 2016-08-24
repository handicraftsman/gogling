/* main.go
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
