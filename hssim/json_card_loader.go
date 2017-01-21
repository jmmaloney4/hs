// Copyright (C) 2017 Jack Maloney. All Rights Reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package hssim

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"strings"
)

type JsonCard struct {
	ID          string
	Name        string
	Set         string
	Type        string
	PlayerClass string
	Cost        uint
	// Collectable bool
	Mechanics []string
	Text      string
	// Artist string
	// Flavor string

	// Minion specific
	Attack int // Shared with Weapon
	Health int
	Race   string

	//Spell specific

	//Weapon specific
	// Borrows Attack from Minion
	Durability int
}

func (game *Game) LoadCardsFromJsonFile(path string) error {
	// read JSON file
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Initialize datastructure for Unmarshall to fill in
	cards := []JsonCard{}
	json.Unmarshal(file, &cards)

	// Filter out Basic cardset
	basicSet := make([]JsonCard, 0)
	count := 0
	for _, c := range cards {
		if c.Set == "CORE" {
			basicSet = append(basicSet, c)
			if c.Type == "WEAPON" {
				//fmt.Println(c)
				// fmt.Println(c.Text)
			}
			count++
		}
	}

	game.cardIndex = make([]Card, 0)

	for _, c := range basicSet {
		abs := AbstractCard{id: c.ID, name: c.Name, class: ClassFromString(c.PlayerClass), cost: c.Cost, text: strings.Replace(c.Text, "\n", " ", -1)}

		switch c.Type {
		case "MINION":
			minion := BasicMinionCard{abs, c.Attack, c.Health, MinionRaceFromString(c.Race), false}
			for _, m := range c.Mechanics {
				switch m {
				case "TAUNT":
					minion.taunt = true
				}
			}
			// fmt.Println(minion)
			game.cardIndex = append(game.cardIndex, minion)
			// fmt.Println(game.cardIndex)
		case "SPELL":
			spell := BasicSpellCard{abs}
			game.cardIndex = append(game.cardIndex, spell)
		case "WEAPON":
			weapon := BasicWeaponCard{abs, c.Attack, c.Durability}
			game.cardIndex = append(game.cardIndex, weapon)
			// fmt.Println(weapon)
		}
	}

	// fmt.Println("Card Index: ", game.cardIndex)
	/*
		for i, c := range game.cardIndex {
			fmt.Println(i, c)
		}
	*/

	return nil
}
