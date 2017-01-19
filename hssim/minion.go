package hssim

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
