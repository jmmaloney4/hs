package hssim

import (
    "bytes"
    "strconv"
)

type SpellCard interface {
	Card
}

type BasicSpellCard struct {
	AbstractCard
}

func (card BasicSpellCard) String() string {
    var buf bytes.Buffer
    
    buf.WriteString(card.Name())
    buf.WriteString(" (")
    buf.WriteString(strconv.Itoa(int(card.Cost())))
    buf.WriteString(" Mana")
    if card.Text() != "" {
        buf.WriteString(", ")
        buf.WriteString(card.Text())
    }
    
    buf.WriteString(")")
    
    return buf.String()
}
