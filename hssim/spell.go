// Copyright (C) 2017 Jack Maloney. All Rights Reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package hssim

import (
	"bytes"
	"strconv"
)

type SpellCard interface {
	Card
}

type BasicSpellCard struct {
	AbstractCard
}

func (card BasicSpellCard) String() string {
	var buf bytes.Buffer

	buf.WriteString(card.Name())
	buf.WriteString(" (")
	buf.WriteString(strconv.Itoa(int(card.Cost())))
	buf.WriteString(" Mana")
	if card.Text() != "" {
		buf.WriteString(", ")
		buf.WriteString(card.Text())
	}

	buf.WriteString(")")

	return buf.String()
}
