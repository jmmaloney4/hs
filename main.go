package main

import (
	// "fmt"
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

	p0.LoadDeck("deck.csv", game)
	p1.LoadDeck("deck.csv", game)
	//fmt.Println(err)

	game.StartGame()

	//fmt.Println(game)
}
