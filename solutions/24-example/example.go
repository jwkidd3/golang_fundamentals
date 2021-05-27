package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "This is a test of maps and ordering"
	MakeMapAndPrint(str)
}

func MakeMapAndPrint(s string) {
	words := make(map[string]int)

	for _, word := range strings.Split(s, " ") {
		words[word] = len(word)
	}
	for word := range words {
		fmt.Println(word)
	}
}
