package main

import "fmt"

func variadic(words ...string) { // HL
	for i := 0; i < len(words); i++ {
		fmt.Print(words[i])
		if len(words[i]) == 6 {
			fmt.Print(" ***")
		}
		fmt.Println()
	}
}

// START OMIT
func main() {
	words := []string{"local", "locale", "localize"}
	variadic(words...) // HL
}

// END OMIT
