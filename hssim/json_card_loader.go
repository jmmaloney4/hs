package hssim

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
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

func DecodeClass(cls string) Class {
	switch cls {
	case "DRUID":
		return ClassDruid
	case "HUNTER":
		return ClassHunter
	case "MAGE":
		return ClassMage
	case "PALADIN":
		return ClassPaladin
	case "PRIEST":
		return ClassPriest
	case "ROUGE":
		return ClassRouge
	case "SHAMAN":
		return ClassShaman
	case "WARLOCK":
		return ClassWarlock
	case "WARRIOR":
		return ClassWarrior
	default:
		return ClassNeutral
	}
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
			}
			count++
		}
	}

	game.cardIndex = make([]Card, 0)

	for _, c := range basicSet {
		abs := AbstractCard{id: c.ID, name: c.Name, class: DecodeClass(c.PlayerClass), cost: c.Cost, text: c.Text}

		switch c.Type {
		case "MINION":
			minion := BasicMinionCard{abs, c.Attack, c.Health, 0, false}
			switch c.Race {
			case "BEAST":
				minion.race = MinionRaceBeast
			case "DEMON":
				minion.race = MinionRaceDemon
			case "DRAGON":
				minion.race = MinionRaceDragon
			case "MECHANICAL":
				minion.race = MinionRaceMech
			case "MURLOC":
				minion.race = MinionRaceMurloc
			case "PIRATE":
				minion.race = MinionRacePirate
			case "TOTEM":
				minion.race = MinionRaceTotem
			default:
				minion.race = MinionRaceNeutral
			}
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

	return nil
}
