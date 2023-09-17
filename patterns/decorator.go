package patterns

import "log"

type InterfaceSoldier interface {
	Attack() int
	Defense() int
}

type BasicSoldier struct{}

func (b BasicSoldier) Attack() int {
	return 1
}

func (b BasicSoldier) Defense() int {
	return 1
}

func DisplayStats(soldier InterfaceSoldier) {
	log.Printf("Soldier stats: Attack(%d) Defense(%d)", soldier.Attack(), soldier.Defense())
}

// Decorator
type SoldierWithSword struct {
	ISoldier InterfaceSoldier
}

func (s SoldierWithSword) Attack() int {
	return s.ISoldier.Attack() + 10
}

func (s SoldierWithSword) Defense() int {
	return s.ISoldier.Defense() + 10
}

func Decorator() {
	basicSoldier := BasicSoldier{}
	DisplayStats(basicSoldier)

	soldierWithSword := SoldierWithSword{ISoldier: basicSoldier}
	DisplayStats(soldierWithSword)
}
