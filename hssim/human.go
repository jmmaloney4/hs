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

func (player *HumanPlayer) Hand() []Card {
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

func (player *HumanPlayer) ChooseCard(cards []Card) (int, error) {
	for i, c := range cards {
		fmt.Println(i, c)
	}

	rv := -1
	for rv == -1 {
		fmt.Printf("[0-%d]? ", len(cards)-1)
		_, err := fmt.Scanf("%d", &rv)
		if err != nil {
			return -1, err
		}
        if rv < 0 || rv > len(cards) - 1 {
            rv = -1
        }
	}

	return rv, nil
}

func (player *HumanPlayer) ChooseAction(game *Game) (Action, error) {
	fmt.Println("Choose Action:")
	fmt.Println("0 Play a Card")
	fmt.Println("1 Attack with a Minion")
	fmt.Println("2 Attack with your Hero")
	fmt.Println("3 Use your Hero Power")
	fmt.Print("[0-3]? ")

	r := bufio.NewReader(os.Stdin)
	c, err := r.ReadByte()
	if err != nil {
		return Action{0, nil}, err
	}

	switch c {
	case '0':
		i, err := player.ChooseCard(player.Hand())
		if err != nil {
			return Action{0, nil}, err
		}
		fmt.Println("Playing", player.Hand()[i])
	case '1':

	case '2':

	case '3':

	default:
		fmt.Println(c, "is not an option")
		return player.ChooseAction(game)
	}

	return Action{0, nil}, nil
}
