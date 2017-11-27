package main

import (
	"fmt"
	"math/rand"
)

type Generator struct {
	characters  []Character
	allActions  []Action
	flaws       []Flaw
	causalLinks []CausalLink
}

func (this *Generator) resolveFlaw() {
	randomIndex := rand.Intn(len(this.flaws))
	flaw := this.flaws[randomIndex]
	causalLink := flaw.resolve(this.allActions, this.characters)

	var notFound CausalLink
	if *causalLink != notFound {
		this.flaws = append(this.flaws[:randomIndex], this.flaws[randomIndex+1:]...)
		this.flaws = append(this.flaws, causalLink.getFlaws()...)
		this.causalLinks = append(this.causalLinks, *causalLink)
	}

	fmt.Printf("\nFlaw resolved: \n%+v\n\nGenerated:\n%+v\n", flaw, causalLink)
}

func NewGenerator() (new Generator) {
	path := "data"
	//var err error
	characters, endStep := NewCharacterList(path + "/POCL-fodder-Characters.csv")

	actions := NewActionList(path + "/POCL-fodder-Actions.csv")
	// fmt.Printf("%+v\n", actions)
	var causalLinks []CausalLink
	return Generator{characters, actions, endStep.getFlaws(), append(causalLinks, endStep)} //, plans}
}
