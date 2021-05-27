package main

// START OMIT
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

func main() {
	// Variadic functions can be called in the usual way with individual arguments
	variadic("this", "that", "hello", "Golang")
	variadic("C++", "Python")
	variadic()
}

// END OMIT
