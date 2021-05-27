package main

import (
	"fmt"
	"thing" // singleton, i.e., only one object // HL
)

func main() {
	// try singleton "thing"
	fmt.Println(thing.Thing())
	thing.SetThing(-5)
	fmt.Println(thing.Thing())
}
