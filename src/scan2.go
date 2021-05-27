package main

// START OMIT
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // HL
	for scanner.Scan() {                  // HL
		fmt.Println(scanner.Text()) // HL
	} // HL
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

// END OMIT
