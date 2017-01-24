// Copyright (C) 2017 Jack Maloney. All Rights Reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package hssim

type Player interface {
	Deck() *Deck
	Hand() []Card

	GoFirst() bool

	// LoadDeck(csvPath string, game *Game) error

	// Base Mana
	BaseMana() int
	// Mana available right now
	AvailableMana() int
	// Mana locked by overloads last turn
	LockedMana() int
	// Mana to be locked next turn
	OverloadedMana() int

	SpendMana(n int)
	Overload(n int)

	MulliganInitialHand(game *Game, hand []Card) error
	MulliganCard(game *Game, index int) (bool, error)
	MulliganFinalHand(game *Game) error

	EndTurn(game *Game) error

	BeginTurn(game *Game) error
	AddCardToHand(game *Game, card Card) error

	//TakeAction(game *Game) Action
}
