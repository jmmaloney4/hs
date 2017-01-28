// Copyright (C) 2017 Jack Maloney. All Rights Reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package hssim

import (
	"fmt"
)

/*

   |   Player 0    Indexed always from left to right as though you were in the players spots
   |  2   1   0    Index to add to
   |    1   0      Current Index of Cards
   | ---------------
   |   / \ / \
   |   \ / \ /
   | ---------------
   |   / \ / \ / \
   |   \ / \ / \ /
   | ---------------
   |    0   1   2      Current Index of Cards
   |  0   1   2   3    Index to insert to
   |   Player 1
*/

type Board struct {
	p0Side []Card
	p1Side []Card
}

func (board *Board) AddMinion(card Card, index int, side int) error {
	if card.Type() != CardTypeMinion {
		return fmt.Errorf("%s is not a Minion", card.String())
	}

	var s *[]Card
	if side == 0 {
		s = &board.p0Side
	} else if side == 1 {
		s = &board.p1Side
	} else {
		return fmt.Errorf("Side must be either 0 or 1")
	}

	if index > len(*s) {
		return fmt.Errorf("Index out of range")
	}

	ns := make([]Card, len(*s)+1)
	copy(ns, (*s)[:index])
	ns[index] = card
	copy(ns[index+1:], (*s)[index:])

	*s = ns

	return nil
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

func (board *Board) GetSide(side int) ([]Card, error) {
	if side == 0 {
		return board.p0Side, nil
	} else if side == 1 {
		return board.p1Side, nil
	} else {
		return nil, fmt.Errorf("Side must be either 0 or 1")
	}
}

func (board Board) String() string {
	rv := fmt.Sprintf("Player 0:\n")
	for i, c := range board.p0Side {
		rv += fmt.Sprintf("%d %s\n", i, c.String())
	}

	rv += "Player 1:\n"
	for i, c := range board.p1Side {
		rv += fmt.Sprintf("%d %s\n", i, c.String())
	}

	return rv
}
