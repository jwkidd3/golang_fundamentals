package main

// START OMIT
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // use seconds since 1970 to seed random number generator
	r := rand.Float64()          // generate random number between 0 and 1

	switch { // HL
	case r > 0.1:
		fmt.Println("Common case, 90% of the time")
	default:
		fmt.Println("10% of the time")
	} // HL
}

// END OMIT
