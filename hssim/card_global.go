// Copyright (C) 2017 Jack Maloney. All Rights Reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package hssim

import (
	"encoding/json"
	// "fmt"
	"fmt"
	"io/ioutil"
	"strings"
)

var globalCardIndex []Card

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

func LoadGlobalCardIndexFromJsonFile(path string) error {
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

	globalCardIndex = make([]Card, 0, len(basicSet))

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
			globalCardIndex = append(globalCardIndex, minion)
			// fmt.Println(globalCardIndex)
		case "SPELL":
			spell := BasicSpellCard{abs}
			globalCardIndex = append(globalCardIndex, spell)
		case "WEAPON":
			weapon := BasicWeaponCard{abs, c.Attack, c.Durability}
			globalCardIndex = append(globalCardIndex, weapon)
			// fmt.Println(weapon)
		}
	}

	// fmt.Println("Card Index: ", globalCardIndex)
	/*
		for i, c := range globalCardIndex {
			fmt.Println(i, c)
		}
	*/

	return nil
}

func CardFromName(name string) (Card, error) {
	// fmt.Println("Card Index: ", game.cardIndex)

	for _, c := range globalCardIndex {
		// fmt.Println(n, " == ", c.Name())
		if name == c.Name() {
			return c, nil
			// fmt.Println("old: ", &c, "new: ", &newCard)
			break
		}
	}

	return nil, fmt.Errorf("Card %s Not Found", name)
}
