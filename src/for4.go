package main

// START OMIT

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var num int32
	sec := time.Now().Unix() // seconds since 1970...
	rand.Seed(sec)           // seed the random number generator

	for { // HL
		fmt.Print("Infinite loop...")
		if num = rand.Int31n(10); num == 5 {
			fmt.Println("break!")
			break
		}
		fmt.Println(num)
	} // HL
}

// END OMIT
