package main

import "fmt"

var _ int = setup() // happens first

func init() { // happens second // HL
	fmt.Println("init!")
}

func init() { // happens third // HL
	fmt.Println("something else")
}

func setup() int {
	fmt.Println("calling setup()")
	return 1
}

func main() { // happens last
	fmt.Println("main")
}
