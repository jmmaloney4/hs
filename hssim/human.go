// Copyright (C) 2017 Jack Maloney. All Rights Reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package hssim

import (
	"bufio"
	"fmt"
	"os"
)

type HumanPlayer struct {
	hand           []Card
	deck           Deck
	first          bool
	baseMana       int
	lockedMana     int
	overloadedMana int
	availableMana  int
}

func NewHumanPlayer(first bool) *HumanPlayer {
	rv := new(HumanPlayer)
	rv.first = first
	return rv
}

func (player *HumanPlayer) Deck() *Deck {
	return &player.deck
}

func (player *HumanPlayer) SetDeck(d Deck) {
	player.deck = d
}

func (player HumanPlayer) Hand() []Card {
	return player.hand
}

func (player HumanPlayer) GoFirst() bool {
	return player.first
}

func (player HumanPlayer) BaseMana() int {
	return player.baseMana
}

func (player HumanPlayer) AvailableMana() int {
	return player.availableMana
}

func (player HumanPlayer) LockedMana() int {
	return player.lockedMana
}

func (player HumanPlayer) OverloadedMana() int {
	return player.overloadedMana
}

func (player HumanPlayer) SpendMana(n int) {
	player.availableMana -= n
}

func (player HumanPlayer) Overload(n int) {
	player.overloadedMana += n
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

func (player *HumanPlayer) BeginTurn(game *Game) error {
	fmt.Println("Starting Turn", game.Turn()/2, "For Player", PlayerNumHumanReadable(player))
	return nil
}

func (player *HumanPlayer) AddCardToHand(game *Game, card Card) error {
	player.hand = append(player.hand, card)

	fmt.Println("Drew:", card)

	return nil
}
