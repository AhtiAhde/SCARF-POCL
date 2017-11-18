package main

import(
)

type CausalLink struct {
    flawResolution string
    internalFlaws []Flaws
    previousLinks []CausalLink
    executedLink *CausalLink // must be nillable
    internalLinks []CausalLink
    //utility []ActionCharacterPairsOrSomething
}
