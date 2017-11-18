package main

import(
)

type Flaw struct {
    flaw string
    origin CausalLink
}

type CausalLink struct {
    flawResolution Flaw
    action Action
}
