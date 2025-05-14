package main

import (
	"testing"
)

func TestBearCanEatMeat(t *testing.T) {
	b := Bear{
		Name: "polar",
		BearType: "Alaska polar bear",
	}

	canEat := b.Eat("meat")
	if canEat != true {
		t.Errorf("Bear should able to eat meat, expected: %v", true)
	}
}
