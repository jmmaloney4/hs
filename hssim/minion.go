// Copyright (C) 2017 Jack Maloney. All Rights Reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package hssim

import (
	"fmt"
)

type MinionRace int

const (
	MinionRaceNeutral MinionRace = iota
	MinionRaceBeast
	MinionRaceDemon
	MinionRaceDragon
	MinionRaceMech
	MinionRaceMurloc
	MinionRacePirate
	MinionRaceTotem
)

type MinionCard interface {
	Card
	Attact() int
	Health() int
	Race() MinionRace
	//HasBattlecry() bool
	//HasDeathrattle() bool
	//HasCharge() bool
	Taunt() bool
	//HasDivineShield() bool
	//HasStealth() bool
	//HasWindfury() bool
	//IsFrozen() bool
	//IsScilenced() bool
}

type BasicMinionCard struct {
	AbstractCard
	attack int
	health int
	race   MinionRace
	taunt  bool
}

func (card BasicMinionCard) Type() CardType {
	return CardTypeMinion
}

func (card BasicMinionCard) Attack() int {
	return card.attack
}

func (card BasicMinionCard) Health() int {
	return card.health
}

func (card BasicMinionCard) Race() MinionRace {
	return card.race
}

func (card BasicMinionCard) Taunt() bool {
	return card.taunt
}

func (card BasicMinionCard) String() string {
	rv := fmt.Sprintf("%s [%s] (%d Mana, %d/%d", card.Name(), card.ID(), card.Cost(), card.Attack(), card.Health())

	if card.Race() != MinionRaceNeutral {
		rv += ", "
		rv += StringFromMinionRace(card.Race())
	}
	if card.Text() != "" {
		rv += ", "
		rv += card.Text()
	}

	rv += ")"

	return rv
}
