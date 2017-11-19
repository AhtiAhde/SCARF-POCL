package main

import (
	"math/rand"
	"strings"
)

type Flaw struct {
	flaw   string
	origin CausalLink
}

func (this *Flaw) resolve(allActions []Action, characters []Character) (causalLink *CausalLink) {
	flawParts := strings.Split(this.flaw, ":")
	character := flawParts[0]
	character = character
	// TODO, what to do with chars? :D
	trait := strings.Join(flawParts[1:], ":")

	var nextPossibleActions []Action
	for _, action := range allActions {
		if action.isResolverOf(trait) {
			nextPossibleActions = append(nextPossibleActions, action)
		}
	}
	if len(nextPossibleActions) > 0 {
		length := len(nextPossibleActions) - 1
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
