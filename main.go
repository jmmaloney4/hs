package main

import (
	"fmt"
	"github.com/jmmaloney4/hssim/hssim"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(hssim.ClassShaman)

	rand.Seed(time.Now().UTC().UnixNano())

	p0 := hssim.NewHumanPlayer()
	p1 := hssim.NewHumanPlayer()
	game, _ := hssim.NewGame(p0, p1)
    
    _, err := hssim.NewDeck("deck.csv")
    fmt.Println(err)
    
	fmt.Println(game)
    
    hssim.LoadCardsFromJsonFile("cards.json")
}
