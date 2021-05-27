package main

import "fmt"

func main() {
	// To create an empty map, use make
	sbux := map[string]int{
		"tall":   12,
		"grande": 16,
		"venti":  20,
	}
	fmt.Println(sbux)

	key := "grande"
	_, present := sbux[key]
	fmt.Println(key, "present?", present)
}
