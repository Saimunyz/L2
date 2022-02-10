package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

// Strategy
type PowerOfPlayer interface {
	show(*Player)
}

// Player - context
type Player struct {
	name         string
	lvl          int
	strength     int
	agility      int
	intelligence int
	itemLvl      int
	power        PowerOfPlayer
}

func newPlayer(name string, powerAlgo PowerOfPlayer) *Player {
	return &Player{
		name:         name,
		lvl:          1,
		strength:     10,
		agility:      20,
		intelligence: 50,
		itemLvl:      10,
		power:        powerAlgo,
	}
}

func (p *Player) setPowerAlgo(powerAlgo PowerOfPlayer) {
	p.power = powerAlgo
}

func (p *Player) showPower() {
	p.power.show(p)
}

// PowerByAvgStats - show player lvl power
type PowerByAvgOfStats struct {
}

// show player lvl power
func (pw *PowerByAvgOfStats) show(p *Player) {
	power := (p.agility + p.intelligence + p.strength) / 3.0
	fmt.Printf("Power is equal: %d\n", power)
}

// PowerByPlayerLvl - show player lvl power
type PowerByPlayerLvl struct {
}

// show player lvl power
func (pw *PowerByPlayerLvl) show(p *Player) {
	fmt.Printf("Power is equal: %d\n", p.lvl)
}

// PowerByItemLvl - show player lvl power
type PowerbyItemLvl struct {
}

// show player lvl power
func (pw *PowerbyItemLvl) show(p *Player) {
	fmt.Printf("Power is equal: %d\n", p.itemLvl)
}

/*
 *	pwByLvl := &PowerByPlayerLvl{}
 *	pwByStats := &PowerByAvgStats{}
 *	pwByItemLvl := &PowerbyItemLvl{}
 *
 *	p := newPlayer("Ivan", pwByLvl)
 *	p.showPower()
 *
 *
 *	p.setPowerAlgo(pwByStats)
 *	p.showPower()
 *
 *	p.setPowerAlgo(pwByItemLvl)
 *	p.showPower()
 *
 */
