package main

import (
	//"math/rand"
	"fmt"
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
	characters, endStep := NewCharacterList(path + "/POCL-fodder-Characters.csv")
	endStep = endStep
	actions := NewActionList(path + "/POCL-fodder-Actions.csv")
	//var plans []Plan
	fmt.Printf("%+v\n", endStep)
	return Generator{characters, actions, endStep} //, plans}
}
