package main

import (
	"fmt"
)

func main() {
	b := Bear{
		Name: "polar",
		BearType: "Alaska polar bear",
	}

	fmt.Println("The bear: ", b)
}

type Bear struct {
	Name     string
	BearType string
}

func (b Bear) Eat(food string) bool {
	if food != "meat" {
		return false
	}
	return true
}
