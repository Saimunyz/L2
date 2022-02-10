package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

// RandomEvents (Visitor) - holds all methods that will be applied on all live forms
type RandomEvents interface {
	eventForPlayer(p *Player) string
	eventForGoblin(g *Goblin) string
}

// LiveForms interface for accepting new methods without changin struct
type LiveForms interface {
	Accept(event RandomEvents)
}

// Player - holds player's data
type Player struct {
	nickname string
	hp       int
	mana     int
	exp      int
}

// newPlayer - creates new Player instance
func newPlayer(nickname string) *Player {
	return &Player{
		nickname: nickname,
		hp:       100,
		mana:     100,
		exp:      0,
	}
}

// getDamage - apply damage on Player
func (p *Player) GetDamage(dmg int) {
	p.hp -= dmg
	fmt.Printf("Player %s was attacked on %d and now have %d health\n", p.nickname, dmg, p.hp)
}

// Accept for accepting method from outside
func (p *Player) Accept(event RandomEvents) {
	event.eventForPlayer(p)
}

// Goblin - contains goblin data
type Goblin struct {
	damage int
	hp     int
	lvl    int
}

// newGoblin - creates new Goblin instance
func newGoblin(damage int) *Goblin {
	return &Goblin{
		damage: damage,
		hp:     100,
		lvl:    1,
	}
}

// Attack - deal damage to player
func (g *Goblin) Attack(p *Player) {
	fmt.Printf("Goblin has attacked %s on damage %d\n", p.nickname, g.damage)
}

// Accept for accepting method from outside
func (g *Goblin) Accept(event RandomEvents) {
	event.eventForGoblin(g)
}

// HolyPotato - (visitor) emits holy light that heals human wounds and damages goblins
type HolyPotato struct {
	damage int
	heal   int
}

// newHolyPotato - creates new HolyPotato events
func newHolyPotato(dmg, heal int) *HolyPotato {
	return &HolyPotato{
		damage: dmg,
		heal:   heal,
	}
}

// eventForPlayer - apply event on Plyaer
func (h *HolyPotato) eventForPlayer(p *Player) string {
	p.hp += h.heal
	return fmt.Sprintf("Holy potato heals %s on %d healp points", p.nickname, h.heal)
}

// eventForGoblin - apply even on Goblin
func (h *HolyPotato) eventForGoblin(g *Goblin) string {
	g.hp -= h.damage
	return fmt.Sprintf("Holy potato damages goblin on %d points", h.damage)
}
