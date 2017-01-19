package hssim

type WeaponCard interface {
	Card
	Attack() int
	Durability() int
}

type BasicWeaponCard struct {
	AbstractCard
	attack     int
	durability int
}

func (card BasicWeaponCard) Attack() int {
	return card.attack
}

func (card BasicWeaponCard) Durability() int {
	return card.durability
}
