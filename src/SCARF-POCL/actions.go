package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Action struct {
	name            string
	transformations string
	preconditions   string
}

func (this *Action) getPrecondtions() (preconditions []string) {
	for _, precondition := range strings.Split(this.preconditions, ";") {
		preconditions = append(preconditions, precondition)
	}
	return preconditions
}

func (this *Action) isResolverOf(trait string) bool {
	for _, transformation := range strings.Split(this.transformations, ";") {
		transformationParts := strings.Split(transformation, ":")
		transformationTrait := strings.Join(transformationParts[1:], ":")
		if strings.Contains(trait, transformationTrait) {
			return true
		}
	}
	return false
}

/**
 * For now we bind only for two characters; subject and object
 * This is very likely to be refactored
 */
func (this *Action) bindCharacters(character string, trait string, characters []Character) {
	this.resolveTargetCharacter(character, trait)

}

func (this *Action) resolveTargetCharacter(character string, trait string) {
	for _, transformation := range strings.Split(this.transformations, ";") {
		transformationParts := strings.Split(transformation, ":")
		transformationTrait := strings.Join(transformationParts[1:], ":")
		transformationChar := transformationParts[0]

		if trait == transformationTrait {
			this.bindToCharacter(transformationChar, character)
			break
		}
	}
}

func (this *Action) resolveOtherCharacters(characters []Character) {
	for _, transformation := range strings.Split(this.transformations, ";") {
		transformationParts := strings.Split(transformation, ":")
		transformationTrait := strings.Join(transformationParts[1:], ":")
		transformationChar := transformationParts[0]

		for _, character := range characters {
			if character.hasTrait(transformationTrait) {
				this.bindToCharacter(transformationChar, character.getName())
				break
				// How to signal handled transformations and preconditions
			}
		}
	}
}

func (this *Action) bindToCharacter(label string, character string) {
	fmt.Println("Binding Transformations to " + character + " for label " + label)
	this.transformations = strings.Replace(this.transformations, label+":", character+":", -1)
	this.preconditions = strings.Replace(this.preconditions, label+":", character+":", -1)
}

func NewActionList(path string) []Action {
	var actions []Action

	f, err := os.Open(path)
	if err != nil {
		return nil
	}

	defer f.Close() // this needs to be after the err check

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil
	}

	var new_action Action
	for key, value := range lines {
		if key == 0 {
			continue
		}

		new_action = Action{name: value[0], transformations: value[2], preconditions: value[1]}
		actions = append(actions, new_action)

	}

	//fmt.Printf("%v\n", actions)
	return actions
}
