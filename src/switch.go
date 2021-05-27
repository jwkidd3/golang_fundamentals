package main

import "fmt"

// START OMIT
import "time"
import "math/rand"

func main() {
	sec := time.Now().Unix() // get current time as seconds since 1970
	rand.Seed(sec)           // "seed" the random number generator
	i := rand.Int31n(4)      // generate random number between 0 and 3

	switch i {
	case 0:
		fmt.Print("zero...")
	case 1:
		fmt.Print("one...")
	case 2:
		fmt.Print("two...")
	}
	fmt.Println("ok")
}

// END OMIT
