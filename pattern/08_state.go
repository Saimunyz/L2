package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type State interface {
	pressButton()
}

// Computer
type Computer struct {
	currState State
}

func newComputer() *Computer {
	computer := Computer{}

	computer.changeState(&TurnedOff{computer: &computer})

	return &computer
}

// pressButton - do action depending on curr state
func (c *Computer) pressButton() {
	c.currState.pressButton()
}

// changeState - changes curren state
func (c *Computer) changeState(state State) {
	c.currState = state
}

// TurnedOff State
type TurnedOff struct {
	computer *Computer
}

// pressbutton in turned off state
func (t *TurnedOff) pressButton() {
	fmt.Println("Nothing")
	t.computer.changeState(&TurnedOn{computer: t.computer})
}

// Asleep state
type Asleep struct {
	computer *Computer
}

// pressbutton in asleep state
func (a *Asleep) pressButton() {
	fmt.Println("Wakes up the computer")
	a.computer.changeState(&TurnedOn{computer: a.computer})

}

type TurnedOn struct {
	computer *Computer
}

// pressbutton in turned on state
func (t *TurnedOn) pressButton() {
	fmt.Println("do something")
	// long time not doing anything
	t.computer.changeState(&Asleep{computer: t.computer})
}

// computer := newComputer()
// for i := 0; i < 5; i++ {
// 	computer.pressButton()
// }
