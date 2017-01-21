// Copyright (C) 2017 Jack Maloney. All Rights Reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package hssim

import (
	// "fmt"
	"math/rand"
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
	Deck() Deck
	Hand() []Card

	GoFirst() bool

	LoadDeck(csvPath string, game *Game) error

	MulliganInitialHand(game *Game, hand []Card) error
	MulliganCard(game *Game, index int) (bool, error)
    MulliganFinalHand(game *Game) error
    
    EndTurn(game *Game) error
}

type Game struct {
	players   []Player
	cardIndex []Card
}

type Deck struct {
	contents []Card
}

func (deck Deck) Draw() Card {
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

	return rv
}

func (deck Deck) ShuffleIn(c Card) {

}

func (game *Game) GetCardByName(name string) (Card, error) {
	// fmt.Println("Card Index: ", game.cardIndex)

	for _, c := range game.cardIndex {
		// fmt.Println(n, " == ", c.Name())
		if name == c.Name() {
			return c, nil
			// fmt.Println("old: ", &c, "new: ", &newCard)
			break
		} else {
			continue
		}
		// TODO: Error Card Not Found
	}

	return nil, nil
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
