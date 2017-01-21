// Copyright (C) 2017 Jack Maloney. All Rights Reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package hssim

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

type HumanPlayer struct {
	hand  []Card
	deck  Deck
	first bool
}

func (player HumanPlayer) InputType() InputType {
	return InputTypeCommandLine
}

func NewHumanPlayer(first bool) *HumanPlayer {
	rv := new(HumanPlayer)
	rv.first = first
	return rv
}

func (player HumanPlayer) Deck() Deck {
	return player.deck
}

func (player HumanPlayer) Hand() []Card {
	return player.hand
}

func (player HumanPlayer) GoFirst() bool {
	return player.first
}

func PlayerNumHumanReadable(player Player) int {
	if player.GoFirst() {
		return 1
	} else {
		return 2
	}
}

func (player *HumanPlayer) MulliganInitialHand(game *Game, hand []Card) error {
	player.hand = hand
	fmt.Print("Mulligan Player ", PlayerNumHumanReadable(player), ":\n")

	for _, c := range hand {
		fmt.Println(c)
	}

	return nil
}

func (player *HumanPlayer) MulliganCard(game *Game, index int) (bool, error) {
	r := bufio.NewReader(os.Stdin)
    fmt.Print("Mulligan the ", player.Hand()[index].Name(), "? [Y/n]: ")
	rune, _, err := r.ReadRune()
    defer r.ReadLine()
	if err != nil {
		return false, err
	} else if rune == 'y' || rune == 'Y' {
		return true, nil
    } else {
        return false, nil
    }
}

func (player *HumanPlayer) MulliganFinalHand(game *Game) error {
    fmt.Println("Starting Hand:")
    for i, _ := range player.Hand() {
        fmt.Println(player.Hand()[i])
	}
    
    return nil
}

func (player *HumanPlayer) EndTurn(game *Game) error {
    fmt.Print("End Turn")
    r := bufio.NewReader(os.Stdin)
    r.ReadLine()
    fmt.Print("\n")
    return nil
}

func (player *HumanPlayer) LoadDeck(csvPath string, game *Game) error {
	file, err := os.Open(csvPath)
	if err != nil {
		return err
	}

	r := csv.NewReader(file)
	rec, err := r.Read()
	if err != nil {
		return err
	}

	d := make([]Card, 0)

	for _, n := range rec {
		// fmt.Println(n)
		// fmt.Println(game.cardIndex)
		c, err := game.GetCardByName(n)
		if err != nil {
			return err
		}
		d = append(d, c)
	}

	player.deck = Deck{d}

	return nil
}

func (player HumanPlayer) Mulligan(gofirst bool) error {

	var cards int
	if gofirst {
		cards = 3
	} else {
		cards = 4
	}

	player.hand = make([]Card, cards)

	var pn int
	if gofirst {
		pn = 1
	} else {
		pn = 2
	}

	fmt.Println("Player ", pn, " Mulligan:")
	for i, _ := range player.hand {
		player.hand[i] = player.deck.Draw()
		fmt.Println(player.hand[i])
	}

	r := bufio.NewReader(os.Stdin)

	for i, c := range player.hand {
		fmt.Print("Mulligan the ", c.Name(), "? [Y/n]: ")
		rune, _, err := r.ReadRune()
		if err != nil {
			return err
		} else if rune == 'y' || rune == 'Y' {
			nc := player.deck.Draw()
			player.deck.contents = append(player.deck.contents, c)
			player.hand[i] = nc
		}
		r.ReadLine()
	}

	fmt.Println("Final Hand: ")
	for i, _ := range player.hand {
		fmt.Println(player.hand[i])
	}

	fmt.Print("End Turn")
	r.ReadLine()

	fmt.Print("\n")

	return nil
}

func (player HumanPlayer) StartTurn() error {

	return nil
}
