// Copyright (C) 2017 Jack Maloney. All Rights Reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	// "fmt"
	"github.com/jmmaloney4/hssim/hssim"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	p0 := hssim.NewHumanPlayer(true)
	p1 := hssim.NewHumanPlayer(false)
	game, _ := hssim.NewGame(p0, p1)

	game.LoadCardsFromJsonFile("cards.json")

	d, _ := hssim.DeckFromCSV("deck.csv", game)
	p0.SetDeck(d)
	d, _ = hssim.DeckFromCSV("deck.csv", game)
	p1.SetDeck(d)
	//fmt.Println(err)

	game.StartGame()

	//fmt.Println(game)
}
