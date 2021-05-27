package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var num int64
	rand.Seed(time.Now().Unix()) // seed the random number generator

	for num != 5 { // HL
		num = rand.Int63n(15)
		fmt.Println(num)
	} // HL
}
