// Write a function doubler which takes a string argument and
// doubles it, i.e., "Golang" would become "GolangGolang"
// ...once you've done that, make it so that doubler calls a
// second function, doublerHelper which actually does the doubling

package main

import "fmt"

func doubler(s *string) {
	// To call another function, we just pass the pointer by value,
	// i.e., we DO NOT pass a pointer to the pointer. We need the
	// pointer (or a copy of it) to change the value it's pointing
	// to. We are not changing the pointer itself.
	doublerHelper(s)
}

func doublerHelper(s *string) {
	*s += *s // *s = *s + *s
}

func main() {
	str := "Golang"
	doubler(&str)
	fmt.Println(str)
}
