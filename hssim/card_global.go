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

    PlayRequirements map[string]int //`json:"playRequirements,omitempty"`

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

	globalCardIndex = make([]Card, 0)

	for _, c := range cards {
		// Filter down to just the Basic cardset
		if c.Set != "CORE" {
			continue
		}

		abs := AbstractCard{id: c.ID, name: c.Name, class: ClassFromString(c.PlayerClass), cost: c.Cost, text: strings.Replace(c.Text, "\n", " ", -1)}

		switch c.Type {
		case "MINION":
			minion := AbstractMinionCard{abs, c.Attack, c.Health, c.Health, MinionRaceFromString(c.Race)}
			for _, m := range c.Mechanics {
				switch m {
				case "TAUNT":
					//minion.taunt = true
				}
			}
            
            // Create new location for card and get pointer to it
            storage := new(AbstractMinionCard)
            *storage = minion
            
            globalCardIndex = append(globalCardIndex, storage)
		case "SPELL":
			spell := AbstractSpellCard{abs}
            
			globalCardIndex = append(globalCardIndex, spell)
		case "WEAPON":
			weapon := AbstractWeaponCard{abs, c.Attack, c.Durability}
			globalCardIndex = append(globalCardIndex, weapon)
			// fmt.Println(weapon)
		}
	}

	return nil
}

func CardFromName(name string) (Card, error) {
	for _, c := range globalCardIndex {
		if name == c.Name() {
			return c, nil
		}
	}

	return nil, fmt.Errorf("Card %s Not Found", name)
}

func CardFromID(id string) (Card, error) {
	for _, c := range globalCardIndex {
		if id == c.ID() {
			return c, nil
		}
	}

	return nil, fmt.Errorf("Card %s Not Found", id)
}
