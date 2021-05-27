package main

import (
	"fmt"
	"os"      // for os.Args
	"strconv" // for strconv.Atoi
)

func main() {
	// call Atoi on 1st argument (not 0th)
	// ignore errors for now
	num1, _ := strconv.Atoi(os.Args[1])
	// call Atoi on 2nd argument (not 0th)
	// ignore errors for now
	num2, _ := strconv.Atoi(os.Args[2])
	// what happens if the user did not enter integers?
	fmt.Println(num1 + num2)
}
