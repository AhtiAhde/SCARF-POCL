package main

import (
    "math/rand"
)

type Generator struct {
	characters      []Character
	allActions      []Action
	flaws           []Flaws
	plans           []Plan
}

func (this *Generator) resolveFlaw() {
    flaw := flaws[rand.Intn(len(this.flaws)) - 1]
    causalLink := flaw.resolve(this.allActions)
}

func NewGenerator() (new Generator) {
	path := "data"
	//var err error
	characters, flaws, endStep := NewCharacterList(path + "/characters.csv")
	actions := NewActionList(path + "/actions.csv")
	var plans []Plan
	return Generator{characters, actions, flaws, plans}
}
