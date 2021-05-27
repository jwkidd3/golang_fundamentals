package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sec := time.Now().Unix()   // seconds since 1970...
	rand.Seed(sec)             // seed the random number generator
	num := rand.Int31n(10) + 5 // generate random num between 5 and 14

	for ; num > 0; num-- { // HL
		fmt.Println(num)
	} // HL
	fmt.Println("Blast off!")
}
