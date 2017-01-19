package hssim

import (
	"encoding/csv"
	"fmt"
    "os"
    "bufio"
)

type HumanPlayer struct {
	hand []Card
	deck Deck
}

func (player HumanPlayer) InputType() InputType {
	return InputTypeCommandLine
}

func NewHumanPlayer() *HumanPlayer {
	rv := new(HumanPlayer)
	return rv
}

func (player HumanPlayer) Deck() *Deck {
	return &player.deck
}

func (player HumanPlayer) Hand() *[]Card {
	return &player.hand
}

func (player *HumanPlayer) LoadDeck(csvPath string, game *Game) error {
	file, err := os.Open(csvPath)
	if err != nil {
		return err
	}

	r := csv.NewReader(file)
	rec, err := r.Read()
	if err != nil {
		return err
	}
	// fmt.Println(rec)

	d := make([]Card, 0)

	for _, n := range rec {
		// fmt.Println(n)
		// fmt.Println(game.cardIndex)
		c, err := game.GetCardByName(n)
		if err != nil {
			return err
		}
		d = append(d, c)
	}

	player.deck = Deck{d}

	return nil
}

func (player HumanPlayer) Mulligan(gofirst bool) error {

	var cards int
	if gofirst {
		cards = 3
	} else {
		cards = 4
	}

	player.hand = make([]Card, cards)

	var pn int
	if gofirst {
		pn = 1
	} else {
		pn = 2
	}

	fmt.Println("Player ", pn, " Mulligan:")
	for i, _ := range player.hand {
		player.hand[i] = player.deck.Draw()
        fmt.Println(player.hand[i].Name(), "(", player.hand[i].Cost(), "Mana )")
	}
    
    r := bufio.NewReader(os.Stdin)
    
    for i, c := range player.hand {
        fmt.Print("Mulligan the ", c.Name(), "? [Y/n]: ")
        rune, _, err := r.ReadRune()
        if err != nil {
            return err
        } else if rune == 'y' || rune == 'Y' {
            nc := player.deck.Draw()
            // fmt.Println("Drew a(n)", nc.Name(), "instead")
            player.deck.contents = append(player.deck.contents, c)
            player.hand[i] = nc
        }
        r.ReadLine()
    }
    
    fmt.Println("Final Hand: ")
    for i, _ := range player.hand {
        fmt.Println(player.hand[i].Name(), "(", player.hand[i].Cost(), "Mana )")
    }
    
    fmt.Print("End Turn")
    r.ReadLine()
    
    fmt.Print("\n")

    return nil
}
