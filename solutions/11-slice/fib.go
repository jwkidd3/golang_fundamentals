// Write a function `fib` that returns a slice containing the first n
// Fibonacci numbers
//
// The Fibonacci sequence is 1, 1, 2, 3, 5, 8...
//
// Each number is the sum of the previous Fibonacci numbers
//
// e.g., `fib(6)` should return an int slice of size 6 containing
// 1, 1, 2, 3, 5, 8

package main

import "fmt"

// For convenience, we'll define n < 2 as an error. That is, the
// caller can't ask for 1 (or 0) Fibonacci number(s).
func fib(n int) []int {
	if n < 2 { // error, return empty slice
		return make([]int, 0)
	}
	nums := make([]int, n)
	nums[0], nums[1] = 1, 1

	for i := 2; i < n; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}
	return nums
}

func main() {
	var num int

	fmt.Print("How many Fibonacci numbers? ")
	fmt.Scanln(&num)
	fmt.Println(fib(num))
}
