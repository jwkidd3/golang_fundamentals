package main

import "fmt"

// START OMIT
func main() {
	s := make([]string, 3)
	s[0], s[1], s[2] = "zero", "one", "two"
	fmt.Println("initial slice:", s)
	copy_of_s := make([]string, len(s)) // HL
	copy(copy_of_s, s)                  // HL
	copy_of_s[0] = "ZERO!"
	fmt.Printf("%v\n%v", copy_of_s, s)
}

// END OMIT
