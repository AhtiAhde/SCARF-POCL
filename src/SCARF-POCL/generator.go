package main

import(
)

type Generator struct {
    characters  []Character
    actions     []Action
    flaws       []Flaws
    causalLinks []CausalLink
}

func NewGenerator() (new Generator) {
    path := "/data"
	//var err error
	characters, flaws := NewCharacterList(path + "characters.csv")
	actions := NewActionList(path + "actions.csv")
	var causalLinks []CausalLink
	return Generator{characters, actions, flaws, causalLinks}
}
