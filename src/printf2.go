package main

import "fmt"

func main() {
	var f float64 = 38730204.3832
	fmt.Printf("%v\n", f)   // print f as a "value"
	fmt.Printf("%f\n", f)   // print f as a float
	fmt.Printf("%.2f\n", f) // restrict f to 2 places after decimal
	var i int = 13
	fmt.Printf("|%5d|\n", i)  // print i 5 spaces wide
	fmt.Printf("|%-5d|\n", i) // how is this different?
	var b bool = false
	fmt.Printf("%t\n", b)
	fmt.Printf("%T %T %T\n", f, i, b)
}
