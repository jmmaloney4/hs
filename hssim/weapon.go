// Copyright (C) 2017 Jack Maloney. All Rights Reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package hssim

import (
	"fmt"
)

type WeaponCard interface {
	Card
	Attack() int
	Durability() int
}

type AbstractWeaponCard struct {
	AbstractCard
	attack     int
	durability int
}

func (card AbstractWeaponCard) Attack() int {
	return card.attack
}

func (card AbstractWeaponCard) Durability() int {
	return card.durability
}

func (card AbstractWeaponCard) String() string {
	rv := fmt.Sprintf("%s [%s] (%d Mana, %d/%d", card.Name(), card.ID(), card.Cost(), card.Attack(), card.Durability())

	if card.Text() != "" {
		rv += ", "
		rv += card.Text()
	}

	rv += ")"

	return rv
}
