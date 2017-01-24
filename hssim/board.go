// Copyright (C) 2017 Jack Maloney. All Rights Reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package hssim

import (
	"fmt"
)

type Board struct {
	p0Side []MinionCard
	p1Side []MinionCard
}

func (board *Board) AddMinion(card Card, index int, side int) {

}

func (board *Board) GetMinionCount(side int) (int, error) {
	if side == 0 {
		return len(board.p0Side), nil
	} else if side == 1 {
		return len(board.p1Side), nil
	} else {
		return 0, fmt.Errorf("Side must be either 0 or 1")
	}
}
