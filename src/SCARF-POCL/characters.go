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

func NewCharacterList(file string) (characters []Character, flaws []Flaws, endStep CausalLink) {
	rawCSVdata, err := ReadCSV(file)
	if err != nil {
		log.Fatal(err)
	}

	for _, each := range rawCSVdata {
		characters = append(characters, Character{name: each[0], intentions: []string{each[1]}, states: []string{each[2]}})
	}
	
	// TODO: Create a list of characters with traits (fields Name, Start)
	// TODO: Create a CausalLink (end step) and flaws, with those traits and end step as origin

	return characters, flaws
}
