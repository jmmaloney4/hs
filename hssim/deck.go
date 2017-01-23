// Copyright (C) 2017 Jack Maloney. All Rights Reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package hssim

import (
	"encoding/csv"
	"math/rand"
	"os"
)

type Deck struct {
	contents []Card
}

func (deck *Deck) Draw() Card {
	r := rand.Int()
	if r < 0 {
		r = r * -1
	}

	r %= len(deck.contents)
	rv := deck.contents[r]

	nc := make([]Card, 0, len(deck.contents)-1)
	for i, c := range deck.contents {
		if i != r {
			nc = append(nc, c)
		}
	}

	deck.contents = nc

	return rv
}

func (deck *Deck) ShuffleIn(c Card) {
	deck.contents = append(deck.contents, c)
}

func DeckFromCSV(csvPath string, game *Game) (Deck, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return Deck{nil}, err
	}

	r := csv.NewReader(file)
	rec, err := r.Read()
	if err != nil {
		return Deck{nil}, err
	}

	d := make([]Card, 0)

	for _, n := range rec {
		// fmt.Println(n)
		// fmt.Println(game.cardIndex)
		c, err := CardFromName(n)
		if err != nil {
			return Deck{nil}, err
		}
		d = append(d, c)
	}

	return Deck{d}, nil
}
