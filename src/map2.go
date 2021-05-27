package main

import "fmt"

func main() {
	// To create an empty map, use make
	sbux := make(map[string]int)
	// Set key/value pairs
	sbux["tall"] = 12
	sbux["grande"] = 16
	sbux["venti"] = 20
	fmt.Println(sbux)
	fmt.Println("len:", len(sbux)) // HL
	delete(sbux, "grande")         // HL
	fmt.Println(sbux)
}
