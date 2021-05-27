package main

import "fmt"

func main() {
	// START OMIT
	num := 1111

	switch {
	case num > 1000:
		fmt.Println("greater than 1000")
		fallthrough // go to next statement // HL

	case num > 100:
		fmt.Println("greater than 100") // (this is next statment) // HL
		fallthrough                     // go to next statement // HL

	case num > 10:
		fmt.Println("greater than 10") // (this is next statement) // HL

	default:
		fmt.Println("less than 10")
	}
	// END OMIT
}
