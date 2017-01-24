// Copyright (C) 2017 Jack Maloney. All Rights Reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package hssim

import (
	"encoding/csv"
	//"fmt"
	"math/rand"
	"os"
)

type Deck struct {
	cards []Card
}

func (deck *Deck) Draw() Card {
	if deck.Size() <= 0 {
		return nil
	}

	r := rand.Int()
	if r < 0 {
		r = r * -1
	}

	r %= len(deck.cards)
	rv := deck.cards[r]

	copy(deck.cards[r:len(deck.cards)-1], deck.cards[r+1:])
	deck.cards = deck.cards[:len(deck.cards)-1]

	/*
		    fmt.Println("LEN:", len(deck.cards))

		    for i, c := range deck.cards {
				fmt.Println(i, c)
			}
	*/

	return rv
}

func (deck Deck) Size() int {
	return len(deck.cards)
}

func (deck *Deck) ShuffleIn(c Card) {
	deck.cards = append(deck.cards, c)
}

func DeckFromCSV(csvPath string, game *Game) (Deck, error) {
	// Open File
	file, err := os.Open(csvPath)
	if err != nil {
		return Deck{nil}, err
	}

	// Load File and Parse CSV
	r := csv.NewReader(file)
	rec, err := r.Read()
	if err != nil {
		return Deck{nil}, err
	}

	cards := make([]Card, 0, len(rec))

	for _, n := range rec {
		c, err := CardFromName(n)
		if err != nil {
			return Deck{nil}, err
		}
		cards = append(cards, c)
	}

	/*
		fmt.Println("DECK Initial:")

		for i, c := range d.cards {
			fmt.Println(i, c)
		}
	*/
	return Deck{cards}, nil
}
