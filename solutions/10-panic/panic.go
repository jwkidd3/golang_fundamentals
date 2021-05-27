package main

import (
	"fmt"
	"math/rand"
)

func input() {
	val := 0
	isSeeded := false

	// A deferred function which performs a "cleanup action" if
	// the panic is due to a negative number. (Remember that in
	// practice we wouldn't panic for something as silly as this.)
	// If the panic is something else, we will re-panic and let
	// a function up the chain handle it (or not).
	defer func() {
		if r := recover(); r == "negative number" {
			fmt.Println("cleanup action")
		} else {
			panic(r)
		}
	}()

	for {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)

		if !isSeeded {
			rand.Seed(int64(val))
			isSeeded = true
		}
		if rand.Float64() > 0.9 {
			panic("random")
		}

		switch {
		case val < 0:
			panic("negative number")
		case val == 0:
			fmt.Println("0 detected")
			return
		default:
			fmt.Println("processing", val)
		}
	}
}

func main() {
	input()
}
