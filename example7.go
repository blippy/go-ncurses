// seehuhn.de/go/goncurses - a Go-wrapper for the ncurses library
// Copyright (C) 2018  Jochen Voss <voss@seehuhn.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// +build ignore

package main

import (
	//"fmt"
	"time"

	//"seehuhn.de/go/goncurses"
	gc "github.com/blippy/goncurses"
)

var scr *gc.Window

func main() {
	scr = gc.Init()
	defer gc.EndWin()
	scr.Print("Hello world")
	scr.Refresh()
	time.Sleep(5 * time.Second)

	/*
	table := []struct {
		Name string
		Code goncurses.AttrType
	}{
		{"normal", goncurses.AttrNormal},
		{"standout", goncurses.AttrStandout},
		{"underline", goncurses.AttrUnderline},
		{"reverse", goncurses.AttrReverse},
		{"blink", goncurses.AttrBlink},
		{"dim", goncurses.AttrDim},
		{"bold", goncurses.AttrBold},
		{"protect", goncurses.AttrProtect},
		{"invis", goncurses.AttrInvis},
		{"altcharset", goncurses.AttrAltcharset},
	}

	win.AddStr("\n")
	for _, entry := range table {
		win.AddStr(fmt.Sprintf("%12s ", entry.Name))
		win.AttrOn(entry.Code)
		win.AddStr("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
		win.AttrOff(entry.Code)
		win.AddStr("\n")
	}
	win.Refresh()

	win.AddStr("\npress any key\n")
	win.GetCh()
	*/
}
