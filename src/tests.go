/* tests.go
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

/** Not Yet Implemented
import (
	"log"
	"os"
)

func tTestNotify(iName string) {
	log.Printf("# Test: running test \"%s\"", iName)
}
func tTestFail(iName string) {
	log.Printf("# Test: test \"%s\" just failed!", iName)
}

func tRunTests() {
	if *sTestName != "none" { // If we are want to test preprocessor
		switch *sTestName {
		case "all": // If we want to run ALL tests
			log.Println("# Test: Running all tests")
			t0echo()

		case "0_echo": // Test "echo" func
			t0echo()

		} // End of switch

		os.Exit(0) // Exit successfully
	}
}

func t0echo() { // Test "echo" func
	tTestNotify("0_echo")
	//if preMain("<?! echo('test') !?>") != "test" { // Run test. If failed - crash
	//	tTestFail("0_echo")
	//	os.Exit(1)
	//}
}
**/
