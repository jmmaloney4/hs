package main

import (
	"fmt"
	"github.com/jmmaloney4/hssim/hssim"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	p0 := hssim.NewHumanPlayer()
	p1 := hssim.NewHumanPlayer()
	game, _ := hssim.NewGame(p0, p1)

    game.LoadCardsFromJsonFile("cards.json")

	game.GetCardByName("")

	_, err := game.LoadDeck("deck.csv")
	fmt.Println(err)

	fmt.Println(game)
}
