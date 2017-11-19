package main

import (
//"math/rand"
)

type Generator struct {
	characters []Character
	allActions []Action
	endStep    CausalLink
	//plans           []Plan
}

func (this *Generator) resolveFlaw() {
	//flaw := this.flaws[rand.Intn(len(this.flaws))-1]
	//causalLink := flaw.resolve(this.allActions, this.characters)
	//causalLink = causalLink
}

func NewGenerator() (new Generator) {
	path := "data"
	//var err error
	characters, endStep := NewCharacterList(path + "/characters.csv")
	endStep = endStep
	actions := NewActionList(path + "/actions.csv")
	//var plans []Plan
	return Generator{characters, actions, endStep} //, plans}
}
