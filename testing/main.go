package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

var p = person{}
var p2 = person{"Bob", 20}
var p3 = person{name: "Alice", age: 30}
var p4 = &person{}
var p5 = &person{"Fred", 50}
var p6 = &person{name: "George", age: 60}

type power struct {
	attack  int
	defense int
}

type location struct {
	x float32
	y float32
	z float32
}

type nonPlayerCharacter struct {
	name  string
	speed int
	hp    int
	power power
	loc   location
}

type attacker struct {
	attackPower int
	damagePower int
}

type sword struct {
	attacker
	twohanded bool
}

type gun struct {
	attacker
	bulletsremaining int
}

type weapon interface {
	Wield() bool
}

func (s sword) Wield() bool {
	fmt.Println("You've wielded a sword!")
	return true
}

func (g gun) Wield() bool {
	fmt.Println("You've wielded a gun!")
	return true
}

func wielder(w weapon) bool {
	fmt.Println("Wielding weapon...")
	return w.Wield()
}

func main() {
	fmt.Println("structs...")

	demon := nonPlayerCharacter{
		name:  "Alfred",
		speed: 21,
		hp:    1000,
		power: power{attack: 75, defense: 50},
		loc:   location{x: 1075.123, y: 521.123, z: 211.231},
	}

	fmt.Println(demon)

	anotherDemon := nonPlayerCharacter{
		name:  "Beelzebub",
		speed: 30,
		hp:    5000,
		power: power{attack: 10, defense: 10},
		loc:   location{x: 32.03, y: 72.45, z: 65.231},
	}
	fmt.Println(anotherDemon)
}
