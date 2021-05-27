package main

import "fmt"

// START OMIT
import "errors"

var CountErr = errors.New("Invalid count")                 // HL
var EmptyErr = errors.New("Can't replicate empty string!") // HL
// END OMIT

// Like Python * for strings, i.e., string replication
func stringstar(str string, count int) (string, error) {
	result := ""
	if count < 1 {
		return "", CountErr // HL
	}
	if str == "" {
		return "", EmptyErr // HL
	}
	for i := 1; i <= count; i++ {
		result += str
	}
	return result, nil // HL
}

func main() {
	res, err := stringstar("Golang", 4) // HL
	fmt.Println("result:", res, "\nerror:", err)
}

// END2 OMIT
