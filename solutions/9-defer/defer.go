package main

import "fmt"

func main() {
	val := 1

	// Deferred cleanup action...
	defer func() {
		fmt.Println("cleanup action")
	}()

	for val != 0 {
		fmt.Print("Enter int: ")
		fmt.Scanf("%d", &val)
		// error condition
		if val < 0 {
			fmt.Println("Error: negative")
			break
		}
	}
	fmt.Println("exiting")
}
