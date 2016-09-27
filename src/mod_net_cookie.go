/* mod_net_cookie.go
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
	"net/http"
	"time"

	lua "github.com/yuin/gopher-lua"
)

func mNetCookieLoader(lLua *lua.LState) int {
	mod := lLua.SetFuncs(lLua.NewTable(), mNetCookieExports) // Register Funtions
	lLua.Push(mod)                                           // Return Module

	return 1
}

var mNetCookieExports = map[string]lua.LGFunction{ // Here we are storing functions
	"get": mNetCookieGet,
	"set": mNetCookieSet,
	"del": mNetCookieDel,
}

func mNetCookieGet(iLua *lua.LState) int {
	iName := iLua.ToString(1)
	lCookie, err := gRequest.Cookie(iName)
	if checkRuntimeErr(lLuaP, err) {
		iLua.RaiseError("Cannot get cookie: %s", iName)
	} else {
		iLua.Push(lua.LString(lCookie.Value))
	}
	return 1
}

func mNetCookieSet(iLua *lua.LState) int {
	iName := iLua.ToString(1)
	iValue := iLua.ToString(2)

	iPath := iLua.ToString(3)
	iDomain := iLua.ToString(4)

	iExpire := iLua.ToInt(5)

	lExpire := time.Now().AddDate(0, 0, iExpire)
	lCookie := http.Cookie{
		Name:    iName,
		Value:   iValue,
		Path:    iPath,
		Domain:  iDomain,
		Expires: lExpire,
	}

	http.SetCookie(gWriter, &lCookie)

	return 0
}

func mNetCookieDel(iLua *lua.LState) int {
	iName := iLua.ToString(1)

	lCookie := http.Cookie{
		Name:    iName,
		Value:   "",
		Path:    "/",
		Expires: time.Date(1970, 01, 01, 01, 01, 01, 01, time.UTC),
	}
	http.SetCookie(gWriter, &lCookie)

	return 0
}
