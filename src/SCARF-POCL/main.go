package main

import (
	"fmt"
	//"os"
	//"strings"
)

func main() {
    //os.Args[1]
	fmt.Println("Hello World!")

    generator := NewGenerator()
    generator = generator
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
