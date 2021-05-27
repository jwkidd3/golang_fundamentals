package main

import "fmt"

func product(terms ...int) int {
	result := 1
	for _, term := range terms {
		result *= term
	}
	return result
}

func main() {
	tests := [][]int{
		{1, 2, 3},
		{4, 5},
		{},
	}
	for _, slice := range tests {
		fmt.Println(slice, product(slice...))
	}
}
