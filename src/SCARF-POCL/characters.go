package main

import (
	"encoding/csv"
	//"fmt"
	"log"
	"os"
	"strings"
)

type Character struct {
	name string
	// intentions []string
	traits string
}

func (this *Character) hasTrait(testTrait string) bool {
	for _, charTrait := range strings.Split(this.traits, ";") {
		if charTrait == testTrait {
			return true
		}
	}
	return false
}

func (this *Character) getName() string {
	return this.name
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

/*func (this *Character) consentCheck(action Action) bool {
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
*/

// Generates an error of traits
func getTraits(rawTraits string) []string {
	// There are spaces in csv file
	// Get rid of them
	var traits []string
	for _, each := range strings.Split(strings.Trim(rawTraits, " "), ";") {
		traits = append(traits, each)
	}
	return traits
}

func NewCharacterList(file string) (characters []Character, endStep CausalLink) {
	rawCSVdata, err := ReadCSV(file)
	if err != nil {
		log.Fatal(err)
	}

	var endPreconditions []string
	for key, each := range rawCSVdata {
		if key == 0 {
			continue
		}
		characters = append(characters, Character{name: each[0], traits: each[1]})
		if each[2] != "" {
			endPreconditions = append(endPreconditions, each[0]+":"+each[2])
		}
	}

	keyframeAction := Action{"THE END", "", strings.Join(endPreconditions, ";")}
	endStep = CausalLink{nil, keyframeAction}
	return characters, endStep
}
