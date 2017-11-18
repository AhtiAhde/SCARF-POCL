package main

import (
	"encoding/csv"
	//"fmt"
	"log"
	"os"
	"strings"
)

type Character struct {
	name       string
	// intentions []string
	trait     []string
}

func ReadCSV(filepath string) ([][]string, error) {
	csvfile, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	fields, err := reader.ReadAll()

	return fields, nil
}

func (this *Character) consentCheck(action Action) bool {
	for _, precondition := range strings.Split(action.preconditions, ";") {
		// @TODO Add management of AND / OR logic

		// Precondition string has array of conditions separated by ";"
		// One condition follows a protocol of character:state
		// While this might be premature decision, it might turn out to be helpful
		requirement := strings.Split(precondition, ":")[1]
		if this.hasState(requirement) {
			return true
		}
	}
	return false
}

func (this *Character) hasState(required string) bool {
	for _, state := range this.states {
		if state == required {
			return true
		}
	}
	return false
}

// Generates an error of traits
func getTraits(rawTraits string) []string  {
	// There are spaces in csv file
	// Get rid of them
	var traits []string
	for _,each := range strings.Split(strings.Trim(rawTraits, " "), ";") {
		traits = append(traits, each)
	}
	return traits
}

func NewCharacterList(file string) (characters []Character, flaws []Flaw) {
	rawCSVdata, err := ReadCSV(file)
	if err != nil {
		log.Fatal(err)
	}

	for _, each := range rawCSVdata {
		currentTraits := getTraits(each[1])
		character := Character{name: each[0], trait: currentTraits}
		desiredTraits := getTraits(each[2])

		characters = append(characters, character)

		// TODO: This solution is not good -> change
		// TODO: What if trait exists for final state, but not represented in initial state?
		for _, desiredTrait := range desiredTraits{
			desiredTraitInfo := strings.Split(desiredTrait, ":")
			for _,currentTrait := range currentTraits{
				currentTraitInfo := strings.Split(currentTrait, ":")
				if (desiredTraitInfo[0] == currentTraitInfo[0] && desiredTraitInfo[1] != currentTraitInfo[1]){
					flaws = append(flaws, Flaw{flaw: desiredTraitInfo})
				}
			}
		}
	}

	// TODO: Create a CausalLink (end step) and flaws, with those traits and end step as origin

	return characters, flaws
}
