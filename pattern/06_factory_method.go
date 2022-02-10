package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

// Class - an abstract interface for mmo ingame character class
type Class interface {
	getClassName() string
}

// Assasin - mmoClass
type Assassin struct {
	className string
	charName  string
}

// newAssasin - constructor
func newAssasin(name string) *Assassin {
	return &Assassin{
		className: "Assassin",
		charName:  name,
	}
}

// getClassName - getter for struct name
func (a *Assassin) getClassName() string {
	return a.className
}

// Warrior - mmoClass
type Warrior struct {
	className string
	charName  string
}

// newWarrior - constructor
func newWarrior(name string) *Warrior {
	return &Warrior{
		className: "Warrior",
		charName:  name,
	}
}

// getClassName - getter for struct name
func (w *Warrior) getClassName() string {
	return w.className
}

// getClass - Factory method for creating new struct
func getClass(className, charName string) (Class, error) {
	switch className {
	case "Warrior":
		return newWarrior(charName), nil
	case "Assassin":
		return newAssasin(charName), nil
	}
	return nil, fmt.Errorf("Wrong class name passed")
}

/*
 *	warrior, _ := getClass("Warrior", "Some name")
 *	assassin, _ := getClass("Assassin", "new Name")
 *
 *	fmt.Printf("Class name: %s\n", warrior.className)
 *	fmt.Printf("Class name: %s\n", assassin.className)
 */
