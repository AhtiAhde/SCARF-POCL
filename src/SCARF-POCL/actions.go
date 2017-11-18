package main

import(
    "strings"
    "os"  
    "encoding/csv"
    //"fmt"  
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
    for _, value := range lines {
        new_action = Action{name: value[0], transformation: value[1], preconditions: value[2], input: value[3]}
        actions = append(actions, new_action)

    }

    //fmt.Printf("%v\n", actions)
    return actions
}
