package hssim

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
)

type InputType int
type CardType int
type Class int
type MinionRace int

const (
	HAND_SIZE_DEFAULT int = 10
	DECK_SIZE_DEFAULT int = 30

	InputTypeCommandLine InputType = iota
	InputTypeBot

	CardTypeMinion CardType = iota
	CardTypeSpell
	CardTypeWeapon

	ClassNeutral Class = iota
	ClassDruid
	ClassHunter
	ClassMage
	ClassPaladin
	ClassPriest
	ClassRouge
	ClassShaman
	ClassWarlock
	ClassWarrior

	MinionRaceNeutral MinionRace = iota
	MinionRaceBeast
	MinionRaceDemon
	MinionRaceDragon
	MinionRaceMech
	MinionRaceMurloc
	MinionRacePirate
	MinionRaceTotem
)

type Player interface {
	InputType() InputType
}

type Game struct {
	players   []Player
	cardIndex []Card
}

type Deck struct {
	contents []Card
}

func (deck Deck) Draw() Card {
	r := rand.Int()
	if r < 0 {
		r = r * -1
	}

	r %= len(deck.contents)
	rv := deck.contents[r]

	newContents := make([]Card, len(deck.contents)-1)
	copy(newContents, deck.contents)
	deck.contents = newContents

	return rv
}

func (game *Game) LoadDeck(csvPath string) (*Deck, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(file)
	rec, err := r.Read()
	if err != nil {
		return nil, err
	}
	// fmt.Println(rec)

	d := make([]Card, 0)

	for _, n := range rec {
		// fmt.Println(n)
		// fmt.Println(game.cardIndex)
		for _, c := range game.cardIndex {
			// fmt.Println(n, " == ", c.Name())
			if n == c.Name() {
				var newCard Card
				newCard = c
				d = append(d, newCard)
				// fmt.Println("old: ", &c, "new: ", &newCard)
                break
			}
		}
	}
    
	return nil, nil
}

func (game *Game) GetCardByName(name string) (Card, error) {
    fmt.Println("Card Index: ", game.cardIndex)
	return nil, nil
}

func NewGame(p0 Player, p1 Player) (*Game, error) {
	rv := new(Game)
	rv.players = []Player{p0, p1}
	return rv, nil
}
