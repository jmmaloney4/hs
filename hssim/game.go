// Copyright (C) 2017 Jack Maloney. All Rights Reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package hssim

import (
// "fmt"
)

type InputType int
type CardType int
type Class int
type MinionRace int

const (
	HAND_SIZE_DEFAULT int = 10
	DECK_SIZE_DEFAULT int = 30

	InputTypeCommandLine InputType = iota
	InputTypeBot

	CardTypeMinion CardType = iota
	CardTypeSpell
	CardTypeWeapon

	ClassNeutral Class = iota
	ClassDruid
	ClassHunter
	ClassMage
	ClassPaladin
	ClassPriest
	ClassRouge
	ClassShaman
	ClassWarlock
	ClassWarrior

	MinionRaceNeutral MinionRace = iota
	MinionRaceBeast
	MinionRaceDemon
	MinionRaceDragon
	MinionRaceMech
	MinionRaceMurloc
	MinionRacePirate
	MinionRaceTotem
)

type Player interface {
	InputType() InputType
	Deck() *Deck
	SetDeck(d Deck)
	Hand() []Card

	GoFirst() bool

	// LoadDeck(csvPath string, game *Game) error

	// Max Potential Mana
	TotalMana() int
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

	// StartTurn(game *Game, num int) error
}

type Game struct {
	players []Player
}

func (game *Game) StartGame() {
	game.RunMulliganForPlayer(game.players[0])
	game.RunMulliganForPlayer(game.players[1])

}

func (game *Game) RunMulliganForPlayer(player Player) error {
	var cards int
	if player.GoFirst() {
		cards = 3
	} else {
		cards = 4
	}

	h := make([]Card, cards)
	for i, _ := range h {
		h[i] = player.Deck().Draw()
	}

	player.MulliganInitialHand(game, h)

	for i, c := range player.Hand() {
		b, _ := player.MulliganCard(game, i)
		if b {
			player.Hand()[i] = player.Deck().Draw()
			player.Deck().ShuffleIn(c)
		}
	}

	player.MulliganFinalHand(game)
	player.EndTurn(game)

	return nil
}

func NewGame(p0 Player, p1 Player) (*Game, error) {
	rv := new(Game)
	rv.players = []Player{p0, p1}
	return rv, nil
}
