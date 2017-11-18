package main

import (
	"fmt",
	"os",
	"strings"
)

// TODO: This might be wrong
func (this *Action) ResolvesGoal(characterTrait string) bool {
    
}

type Action struct {
    name string
    transformation string
    preconditions string
    input string
}

func (this *Action) hasConsent(character Character) bool {
    for _, precondition := range strings.Split(this.preconditions, ";") {
        condition := strings.Split(precondition, ":")[1]
        if character.hasState(condition) {
            return true
        }
}

type CausalLink struct {
    flawResolution string
    internalFlaws Flaws
    previousLinks [CausalLink]
    executedLink *CausalLink // must be nillable
    internalLinks [CausalLink]
    utility [ActionCharacterPairsOrSomething]
}

type Character struct {
    name string
    intentions [string]
    states [string]
}

func (this *Character) consentCheck(action Action) {
    for _, precondition := range strings.Split(action.preconditions, ";") {
        // @TODO Add management of AND / OR logic
        
        // Precondition string has array of conditions separated by ";"
        // One condition follows a protocol of character:state
        // While this might be premature decision, it might turn out to be helpful
        requirement := strings.Split(precondition, ":")[1]
        for _, state := range this.states {
            if state == requirement {
                return true
            }
        }
    }
    return false
}

func (this *Character) hasState(state string) {
    for _, state := range this.states {
        if state == requirement {
            return true
        }
    }
    return false
}

type Generator struct {
    characters  [Character]
    actions     [Action]
    flaws       [Flaws]
    causalLinks [CausalLink]
}

func main() {
    os.Args[1]
	fmt.Println("Hello World!")

    	
	// Initialization
	// -Load characters, keyframe A and B, actions
	// -Parse initial flaw and efficient searchable action map
	
	// Start Iterating flaw elimination until length of initial flaw is zero
	// or some other limit is exceeded
	
	// Flaw resolution creates a CausalLink, which resolves the character-state
	// pair, but also creates a new flaw, if the Action has unfulfilled
	// preconditions; if Actions precondition is fulfilled, the CausalLink can
	// be considered as "executableLink" (or "executedLink" if we don't want to
	// consider more alternatives after finding one executable path)
	
	// A utility function should be created for assessing the value of drama
	// in each "executableCausalLinkPath"; in case dramatic value is contested,
	// the dramatic CausalLinkPath can be enriched with more dramatic
	// CausalLinks; merely by using conflicting intensions and SCARF.
}

func populateCharacters
