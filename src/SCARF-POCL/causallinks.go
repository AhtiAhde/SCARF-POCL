package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Flaw struct {
	flaw   string
	origin *CausalLink
}

func (this *Flaw) resolve(allActions []Action, characters []Character) (causalLink *CausalLink) {
	flawParts := strings.Split(this.flaw, ":")
	character := flawParts[0]
	character = character

	trait := strings.Join(flawParts[1:], ":")
	var nextPossibleActions []Action
	for _, action := range allActions {
		if action.isResolverOf(trait) {
			nextPossibleActions = append(nextPossibleActions, action)
		}
	}
	fmt.Printf("%+v\n", nextPossibleActions)
	if len(nextPossibleActions) > 0 {
		length := len(nextPossibleActions)
		fmt.Printf("%v\n", length)
		nextAction := nextPossibleActions[rand.Intn(length)]
		return &CausalLink{this, nextAction}
	}
	var notFound CausalLink
	return &notFound
}

type CausalLink struct {
	flawResolution *Flaw
	action         Action
}

func (this *CausalLink) getFlaws() (flaws []Flaw) {
	for _, flaw := range this.action.getPrecondtions() {
		flaws = append(flaws, Flaw{flaw, this})
	}
	return flaws
}
