// Copyright (C) 2017 Jack Maloney. All Rights Reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"encoding/json"
	// "fmt"
	"fmt"
	"io/ioutil"
	"strings"
    "os"
    "github.com/jmmaloney4/hs/hssim"
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

    PlayRequirements map[string]int

	// Minion specific
	Attack int // Shared with Weapon
	Health int
	Race   string

	//Spell specific

	//Weapon specific
	// Borrows Attack from Minion
	Durability int
}

func main() {
    out, err := os.Create(os.Args[1])
    if err != nil {
        panic(err)
    }

    // read JSON file
	file, err := ioutil.ReadFile("cards.json")
	if err != nil {
        panic(err)
	}

	// Initialize datastructure for Unmarshall to fill in
	cards := []JsonCard{}
	json.Unmarshal(file, &cards)

    out.WriteString(fmt.Sprintf(`
        // ------------------------------
        // GENERATED FILE
        // ------------------------------

        package cards

        import("github.com/jmmaloney4/hs/hssim")

        type NewCardFunc func() Card

        var globalIndexNames map[string]NewCardFunc
        var globalIndexIDs map[string]NewCardFunc
        const TotalCards int = %d

        func init() {
            globalIndexNames := make(map[string]NewCardFunc, TotalCards)
            globalIndexIDs := make(map[string]NewCardFunc, TotalCards)
`, len(cards)))

    for _, c := range cards {
        out.WriteString(fmt.Sprintf("globalIndexNames[\"%s\"] = New%s;\n", strings.Replace(c.Name, "\"", "\\\"", -1), c.ID))
        out.WriteString(fmt.Sprintf("globalIndexIDs[\"%s\"] = New%s;\n\n", c.ID, c.ID))
    }

    out.WriteString("}\n")

    for _, c := range cards {
        switch c.Type {
        case "MINION":
            text := strings.Replace(c.Text, "\n", " ", -1)
            text = strings.Replace(text, "\"", "\\\"", -1)

            name := strings.Replace(c.Name, "\"", "\\\"", -1)

            out.WriteString(fmt.Sprintf("type %s struct { AbstractMinionCard; }\n", c.ID))
            out.WriteString(fmt.Sprintf("func New%s() Card { return &%s{{{\"%s\", \"%s\", Class%s, %d, \"%s\"}, %d, %d, %d, MinionRace%s}} }\n\n", c.ID, c.ID, name, c.ID, "", c.Cost, text, c.Attack, c.Health, c.Health, hssim.StringFromMinionRace(hssim.MinionRaceFromString(c.Race))))

		case "SPELL":
			text := strings.Replace(c.Text, "\n", " ", -1)
			text = strings.Replace(text, "\"", "\\\"", -1)

			name := strings.Replace(c.Name, "\"", "\\\"", -1)

			out.WriteString(fmt.Sprintf("type %s struct { AbstractSpellCard; }\n", c.ID))
			out.WriteString(fmt.Sprintf("func New%s() Card { return &%s{{{\"%s\", \"%s\", Class%s, %d, \"%s\"}}} }\n\n", c.ID, c.ID, name, c.ID, "", c.Cost, text))

		case "WEAPON":
			text := strings.Replace(c.Text, "\n", " ", -1)
			text = strings.Replace(text, "\"", "\\\"", -1)

			name := strings.Replace(c.Name, "\"", "\\\"", -1)

			out.WriteString(fmt.Sprintf("type %s struct { AbstractWeaponCard; }\n", c.ID))
			out.WriteString(fmt.Sprintf("func New%s() Card { return &%s{{{\"%s\", \"%s\", Class%s, %d, \"%s\"}, %d, %d}} }\n\n", c.ID, c.ID, name, c.ID, "", c.Cost, text, c.Attack, c.Durability))

        }

    }
}
