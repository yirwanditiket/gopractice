package main

import (
	"fmt"
)

func main() {
	polarBear := Bear{Name: "snowy", BearType: "polar"}
	yusuf := Human{Name: "Yusuf"}
	cart := Rabbit{Name: "Cart"}

	feed([]Animal{
		polarBear,
		yusuf,
		cart,
	})
}

type Animal interface {
	// Eat accept food parameter, and return bool to tell if the animal can eat the food / not
	Eat(food string) bool
}

type Bear struct {
	Name     string
	BearType string
}

type Human struct {
	Name string
}

type Rabbit struct {
	Name string
}

func (b Bear) Eat(food string) bool {
	if food != "meat" {
		return false
	}
	return true
}

func (h Human) Eat(food string) bool {
	return true
}

func (r Rabbit) Eat(food string) bool {
	if food == "vegetable" {
		return true
	}

	return false
}

func feed(animals []Animal) {
	foods := []string{
		"vegetable",
		"meat",
		"rice",
	}

	for _, food := range foods {
		for _, animal := range animals {
			canEat := animal.Eat(food)
			if canEat {
				fmt.Printf("%v can eat %s\n", animal, food)
			} else {
				fmt.Printf("%v can't eat %s\n", animal, food)
			}
		}
	}
}
