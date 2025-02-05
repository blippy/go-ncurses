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

//import "seehuhn.de/go/goncurses"
import "github.com/blippy/goncurses"

func main() {
	goncurses.Init()
	defer goncurses.EndWin()

	c1 := goncurses.ColorPair(1)
	c1.Init(goncurses.ColorRed, goncurses.ColorYellow)
	c2 := goncurses.ColorPair(2)
	c2.Init(goncurses.ColorGreen, goncurses.ColorBlue)

	w1 := goncurses.NewWin(5, 10, 1, 0)
	w1.SetBackground("X", goncurses.AttrBold, c1)
	w1.Erase()
	w1.Println("some text")
	w1.Refresh()

	w2 := goncurses.NewWin(5, 10, 0, 10)
	w2.SetBackground(".", goncurses.AttrBlink, c2)
	w2.Erase()

	w2.AttrSet(c1.AsAttr())
	w2.Println("some text")
	w2.Refresh()

	w3 := goncurses.NewWin(1, 20, 6, 0)
	w3.Print("press any key ...")
	w3.GetCh()
}
