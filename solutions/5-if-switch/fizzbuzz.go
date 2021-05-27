// Write a function which accepts an integer and returns
//
// "fizz" if the number is divisible by 3
// "buzz" if the number is divisible by 5
// "fizzbuzz" if the number is divisible by BOTH 3 and 5
// otherwise it returns the string version of the number, e.g., "4"
//
// Test your function with these inputs: 3, 5, 15, 4
package main

import "fmt"
import "strconv"

func fizzbuzz(n int) string {
	var fizz, buzz string

	if n%3 == 0 {
		fizz = "fizz"
	}
	if n%5 == 0 {
		buzz = "buzz"
	}
	if fizz+buzz == "" {
		return strconv.Itoa(n)
	}
	return fizz + buzz
}

func fizzbuzz2(n int) string {
	result := ""

	if n%3 == 0 {
		result = "fizz"
	}
	if n%5 == 0 {
		result += "buzz"
	}
	if result == "" {
		return strconv.Itoa(n)
	}
	return result
}

func fizzbuzz3(num int) string {
	switch {
	case num%15 == 0: // divisible by 3 and 5 means divisible by 15
		return "fizzbuzz"
	case num%3 == 0:
		return "fizz"
	case num%5 == 0:
		return "buzz"
	}
	return strconv.Itoa(num)
}

func main() {
	// Forward reference here: we don't know slices yet, but this
	// declares a slice (like a "dynamic" array) of ints that we'll
	// use to test our functions.
	testnums := [...]int{3, 4, 5, 15}

	fmt.Println("input  fizzbuzz  fizzbuzz2  fizzbuzz3")
	fmt.Println("-------------------------------------")
	// Another forward reference: this is how we iterate thru a slice
	for _, num := range testnums {
		fmt.Printf("%5d %9s %10s %10s\n",
			num, fizzbuzz(num), fizzbuzz2(num), fizzbuzz3(num))
	}
}
