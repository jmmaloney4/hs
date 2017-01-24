// Copyright (C) 2017 Jack Maloney. All Rights Reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package hssim

import (
// "fmt"
//"bufio"
//"os"
)

type Class int

const (
	HAND_SIZE_DEFAULT int = 10
	DECK_SIZE_DEFAULT int = 30

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
)

type Game struct {
	players []Player

	// turn 0 is mulligan p1
	// turn 1 is mulligan p2
	// turn 2 is p1 turn 1
	// turn 3 is p2 turn 1
	// etc
	turn int

	board Board
}

func (game *Game) StartGame() {

	game.turn = 0
	game.RunMulliganForPlayer(game.players[0])
	game.turn++
	game.RunMulliganForPlayer(game.players[1])

	game.board.p0Side = make([]Card, 0)
	game.board.p1Side = make([]Card, 0)

	game.turn++
	game.BeginTurnForPlayer(game.players[0], game.Turn())

	/*
	   	game.board = Board{make([]Card, 0), make([]Card, 0)}
	   	c, _ := CardFromName("Bloodfen Raptor")
	   	game.board.AddMinion(c, 0, 0)
	   	fmt.Println(game.board)

	       game.board.AddMinion(c, 0, 1)
	       c, _ = CardFromName("River Crocolisk")
	   	game.board.AddMinion(c, 1, 1)
	   	fmt.Println(game.board)
	*/
}

func (game Game) Turn() int {
	return game.turn
}

func (game *Game) RunMulliganForPlayer(player Player) error {
	var cards int
	if player.GoFirst() {
		cards = 3
	} else {
		cards = 4
	}

	h := make([]Card, 0, cards)
	for i := 0; i < cards; i++ {
		h = append(h, player.Deck().Draw())
	}

	if !player.GoFirst() {
		coin, err := CardFromName("The Coin")
		if err != nil {
			return err
		}
		h = append(h, coin)
	}

	player.MulliganInitialHand(game, h)

	for i, c := range player.Hand() {
		if i >= 4 {
			break
		}
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

func (game *Game) BeginTurnForPlayer(player Player, turn int) error {
	player.BeginTurn(game)

	c := player.Deck().Draw()

	player.AddCardToHand(game, c)

	return nil
}

func NewGame(p0 Player, p1 Player) (*Game, error) {
	rv := new(Game)
	rv.players = []Player{p0, p1}
	return rv, nil
}
