// Copyright (C) 2017 Jack Maloney. All Rights Reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package hssim

import (
	"fmt"
)

type SpellCard interface {
	Card
}

type BasicSpellCard struct {
	AbstractCard
}

func (card BasicSpellCard) String() string {
	rv := fmt.Sprintf("%s [%s] (%d Mana", card.Name(), card.ID(), card.Cost())

	if card.Text() != "" {
		rv += ", "
		rv += card.Text()
	}

	rv += ")"

	return rv
}
