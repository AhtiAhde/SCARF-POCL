package main

import(
)

type Flaw struct {
    flaw string
    origin CausalLink
}

func (this *Flaw) resolve(allActions []Action, characters []Character) {
    flawParts := strings.Split(this.flaw, ":")
    character := flawParts[0]
    trait := strings.Join(flawParts[1:], ":")
    
    for _, action := range allActions {
        action.isResolverOf(trait)
    }
}

type CausalLink struct {
    flawResolution Flaw
    action Action
}
