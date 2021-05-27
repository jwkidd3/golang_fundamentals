package main

import "fmt"

// START OMIT
func main() {
	s := []string{"raspberry", "orange", "guava"}
	fmt.Println("declared slice:", s)
	t := []string{"cherry", "banana"}
	s = append(s, t[0], t[1])
	fmt.Println(s, len(s), cap(s))
}

// END OMIT
