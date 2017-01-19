package hssim

import (
	// "fmt"
	"math/rand"
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
	Deck() *Deck
	Hand() *[]Card

	LoadDeck(csvPath string, game *Game) error

	Mulligan(gofirst bool) error
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

	nc := make([]Card, 0, len(deck.contents)-1)
	for i, c := range deck.contents {
		if i != r {
			nc = append(nc, c)
		}
	}

	return rv
}

func (game *Game) GetCardByName(name string) (Card, error) {
	// fmt.Println("Card Index: ", game.cardIndex)

	for _, c := range game.cardIndex {
		// fmt.Println(n, " == ", c.Name())
		if name == c.Name() {
			return c, nil
			// fmt.Println("old: ", &c, "new: ", &newCard)
			break
		} else {
			continue
		}
		// TODO: Error Card Not Found
	}

	return nil, nil
}

func (game *Game) StartGame() {
	game.players[0].Mulligan(true)
	game.players[1].Mulligan(false)
}

func NewGame(p0 Player, p1 Player) (*Game, error) {
	rv := new(Game)
	rv.players = []Player{p0, p1}
	return rv, nil
}
