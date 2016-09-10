/* config.go
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
	"encoding/json"
	"io/ioutil"
)

var sConf = map[string]string{}

func cMain() {
	lConf.Println("Reading config file")

	lData, err := ioutil.ReadFile("conf.json") // Read our config.
	checkErr(lConf, err)

	lConf.Println("Parsing JSON")

	err = json.Unmarshal(lData, &sConf) // Parse JSON
	checkErr(lConf, err)

	lConf.Println("Finished!") // Done!
	sDone <- true
}
