// Copyright (C) 2017 Jack Maloney. All Rights Reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package hssim

import (
	"fmt"
)

type CardType int

const (
	CardTypeMinion CardType = iota
	CardTypeSpell
	CardTypeWeapon
)

type Card interface {
	fmt.Stringer

	Name() string
	ID() string
	// Set() CardSet
	Type() CardType
	Class() Class
	Cost() uint
	Text() string
}

type AbstractCard struct {
	name  string
	id    string
	class Class
	cost  uint
	text  string
}

func (card AbstractCard) Name() string {
	return card.name
}

func (card AbstractCard) ID() string {
	return card.id
}

func (card AbstractCard) Type() CardType {
	// Just defaults to minion because there's no neutral type, in theory should be overridden
	return CardTypeMinion
}

func (card AbstractCard) Class() Class {
	return card.class
}

func (card AbstractCard) Cost() uint {
	return card.cost
}

func (card AbstractCard) Text() string {
	return card.text
}
