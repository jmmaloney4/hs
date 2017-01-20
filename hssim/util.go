// Copyright (C) 2017 Jack Maloney. All Rights Reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package hssim

func MinionRaceFromString(str string) MinionRace {
	switch str {
	case "BEAST":
		return MinionRaceBeast
	case "DEMON":
		return MinionRaceDemon
	case "DRAGON":
		return MinionRaceDragon
	case "MECHANICAL":
		return MinionRaceMech
	case "MURLOC":
		return MinionRaceMurloc
	case "PIRATE":
		return MinionRacePirate
	case "TOTEM":
		return MinionRaceTotem
	default:
		return MinionRaceNeutral
	}
}

func StringFromMinionRace(race MinionRace) string {
	switch race {
	case MinionRaceBeast:
		return "Beast"
	case MinionRaceDemon:
		return "Demon"
	case MinionRaceDragon:
		return "Dragon"
	case MinionRaceMech:
		return "Mech"
	case MinionRaceMurloc:
		return "Murloc"
	case MinionRacePirate:
		return "Pirate"
	case MinionRaceTotem:
		return "Totem"
	default:
		return "Neutral"
	}
}

func ClassFromString(str string) Class {
	switch str {
	case "DRUID":
		return ClassDruid
	case "HUNTER":
		return ClassHunter
	case "MAGE":
		return ClassMage
	case "PALADIN":
		return ClassPaladin
	case "PRIEST":
		return ClassPriest
	case "ROUGE":
		return ClassRouge
	case "SHAMAN":
		return ClassShaman
	case "WARLOCK":
		return ClassWarlock
	case "WARRIOR":
		return ClassWarrior
	default:
		return ClassNeutral
	}
}
