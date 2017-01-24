// Copyright (C) 2017 Jack Maloney. All Rights Reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	// "fmt"
	"fmt"
	"github.com/jmmaloney4/hs/hssim"
	"math/rand"
	"os"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	p0 := hssim.NewHumanPlayer(true)
	p1 := hssim.NewHumanPlayer(false)
	game, _ := hssim.NewGame(p0, p1)

	hssim.LoadGlobalCardIndexFromJsonFile(os.Args[1])

	d, err := hssim.DeckFromCSV(os.Args[2], game)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	*(p0.Deck()) = d
	d, err = hssim.DeckFromCSV(os.Args[3], game)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	*(p1.Deck()) = d
	//fmt.Println(err)

	game.StartGame()

	//fmt.Println(game)
}
