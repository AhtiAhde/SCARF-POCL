package main

import(
    "strings"    
)

type Action struct {
    name string
    transformation string
    preconditions string
    characters []Character
}

func (this *Action) hasConsent(character Character) bool {
    for _, precondition := range strings.Split(this.preconditions, ";") {
        condition := strings.Split(precondition, ":")[1]
        if character.hasState(condition) {
            return true
        }
    }
    return false
}

func NewActionList(path string) ([]Action) {
    var actions []Action
    return actions
}
