package main

import "fmt"

// START OMIT
func main() {
	var n []int
	z := make([]int, 0)

	fmt.Printf("n is %#v %d %d\n", n, len(n), cap(n))
	if n == nil {
		fmt.Println("n is nil!\n")
	}

	fmt.Printf("z is %#v %d %d\n", z, len(z), cap(z))
	if z == nil {
		fmt.Println("z is nil!")
	} else {
		fmt.Println("z is NOT NOT NOT nil!")
	}
}

// END OMIT
