package main

import(
    "strings"    
)

type Character struct {
    name string
    intentions []string
    states []string
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

func NewCharacterList(file string) (characters []Character, flaws []Flaws) {
    //characters, err := os.Open(file)
    return characters, flaws
}