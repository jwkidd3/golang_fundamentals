package main

import "fmt"

// false indicates error
func Collatz(number int64) bool {
	if number < 1 {
		return false
	}
	// Goal was to write a for loop as a while loop
	// which I basically did here. The pre- and post-
	// conditions refer to count, which is just there
	// to make the output of the program look nice.
	for count := 1; number > 1; count++ {
		fmt.Printf("%5d", number)
		if number%2 == 0 {
			number /= 2
		} else {
			number = number*3 + 1
		}
		// We print a blank line after every 10 numbers
		// that are printed.
		if count%10 == 0 {
			fmt.Println()
		}
	}
	fmt.Println("    1")
	return true
}

func main() {
	Collatz(513)
}
