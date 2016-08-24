/* err.go
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
	"log"
	"os"
)

func checkErr(iPart string, iErr error) {
	if iErr != nil {
		log.Printf("\033[31m# %s: %s \033[0m\n", iPart, iErr.Error())
		os.Exit(1)
	}
}

func checkWarn(iPart string, iErr error) {
	if iErr != nil {
		log.Printf("\033[33m# %s: %s \033[0m\n", iPart, iErr.Error())
	}
}
