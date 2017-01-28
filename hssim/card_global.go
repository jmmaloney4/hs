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
	"sort"
	"strings"
)

var globalCardIndex []Card

type JsonCard struct {
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Set         string   `json:"set,omitempty"`
	Type        string   `json:"type,omitempty"`
	PlayerClass string   `json:"playerClass,omitempty"`
	Cost        uint     `json:"cost,omitempty"`
	Collectible bool     `json:"collectible,omitempty"`
	Rarity      string   `json:"rarity,omitempty"`
	Mechanics   []string `json:"mechanics,omitempty"`
	Text        string   `json:"text,omitempty"`
	Artist      string   `json:"artist,omitempty"`
	Flavor      string   `json:"flavor,omitempty"`

	//TargetingArrowText string `json:"targetingArrowText"`

	PlayRequirements map[string]int `json:"playRequirements,omitempty"`

	// Minion specific
	Attack int    `json:"attack,omitempty"` // Shared with Weapon
	Health int    `json:"health,omitempty"`
	Race   string `json:"race,omitempty"`

	//Spell specific

	//Weapon specific
	// Borrows Attack from Minion
	Durability int `json:"durability"`
}

type JsonCards []JsonCard

func (slice JsonCards) Len() int {
	return len(slice)
}

func (slice JsonCards) Less(i, j int) bool {
	return slice[i].Name < slice[j].Name
}

func (slice JsonCards) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
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

	all := make(JsonCards, 0)

	for _, c := range cards {
		// Filter down to just the Basic cardset
		if c.Set != "CORE" {
			continue
		}

		all = append(all, c)

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

	sort.Sort(all)
	o, err := json.MarshalIndent(all, "", "    ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(o[:]))

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
