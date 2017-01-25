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
	for i := range player.Hand() {
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

func (player *HumanPlayer) ChooseOption(opts []string) (int, error) {
    if len(opts) == 0 {
        return -1, fmt.Errorf("No options to choose from")
    }
    
	for i, s := range opts {
		fmt.Println(i, s)
	}

	rv := -1
	for rv == -1 {
		fmt.Printf("[0-%d]? ", len(opts)-1)
		_, err := fmt.Scanf("%d", &rv)
		if err != nil {
			return -1, err
		}
		if rv < 0 || rv > len(opts)-1 {
			rv = -1
		}
	}

	return rv, nil
}

func (player *HumanPlayer) ChooseCard(cards []Card) (int, error) {
	s := make([]string, 0)
	for _, c := range cards {
		s = append(s, c.String())
	}

	return player.ChooseOption(s)
}

func (player *HumanPlayer) ChooseAction(game *Game) (Action, error) {
	fmt.Println("Choose Action:")
	opts := []string{"Play a Card", "Attack with a Minion", "Attack with your Hero", "Use your Hero Power"}

	i, err := player.ChooseOption(opts)
	if err != nil {
		return Action{0, nil}, err
	}

	switch i {
	case 0:
		i, err := player.ChooseCard(player.Hand())
		if err != nil {
			return Action{0, nil}, err
		}
		fmt.Println("Playing", player.Hand()[i])
	case 1:
        s, err := game.board.GetSide(game.WhosTurn())
        if err != nil {
			return Action{0, nil}, err
		}
        
        i, err := player.ChooseCard(s)
        if err != nil {
			return Action{0, nil}, err
		}
        
		fmt.Println("Attacking with", s[i])

	case 2:

	case 3:

	default:
		fmt.Println("wtf shouldn't be here")
		return player.ChooseAction(game)
	}

	return Action{0, nil}, nil
}
