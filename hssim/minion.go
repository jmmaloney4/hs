package hssim

import (
    "bytes"
    "strconv"
    // "fmt"
    )

type MinionCard interface {
	Card
	Attact() int
	Health() int
	Race() MinionRace
	//HasBattlecry() bool
	//HasDeathrattle() bool
	//HasCharge() bool
	Taunt() bool
	//HasDivineShield() bool
	//HasStealth() bool
	//HasWindfury() bool
	//IsFrozen() bool
	//IsScilenced() bool
}

type BasicMinionCard struct {
	AbstractCard
	attack int
	health int
	race   MinionRace
	taunt  bool
}

func (card BasicMinionCard) Type() CardType {
	return CardTypeMinion
}

func (card BasicMinionCard) Attack() int {
	return card.attack
}

func (card BasicMinionCard) Health() int {
	return card.health
}

func (card BasicMinionCard) Race() MinionRace {
	return card.race
}

func (card BasicMinionCard) Taunt() bool {
	return card.taunt
}

func (card BasicMinionCard) String() string {
	var buf bytes.Buffer
    
    buf.WriteString(card.Name())
    buf.WriteString(" (")
    buf.WriteString(strconv.Itoa(int(card.Cost())))
    buf.WriteString(" Mana, ")
    buf.WriteString(strconv.Itoa(card.Attack()))
    buf.WriteString("/")
    buf.WriteString(strconv.Itoa(card.Health()))
    if card.Race() != MinionRaceNeutral {
        buf.WriteString(", ")
        switch card.Race() {
            case MinionRaceBeast:
            buf.WriteString("Beast")
            case MinionRaceDemon:
            buf.WriteString("Demon")
            case MinionRaceDragon:
            buf.WriteString("Dragon")
            case MinionRaceMech:
            buf.WriteString("Mech")
            case MinionRaceMurloc:
            buf.WriteString("Murloc")
            case MinionRacePirate:
            buf.WriteString("Pirate")
            case MinionRaceTotem:
            buf.WriteString("Totem")
            default:
            panic("Expected A Valid Minion Race")
        }
    }
    if card.Text() != "" {
        buf.WriteString(", ")
        buf.WriteString(card.Text())
    }
    buf.WriteString(")")
    
    return buf.String()
}