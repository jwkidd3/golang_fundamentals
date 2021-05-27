package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	return x * y
}

func div(x, y int) int {
	return x / y
}

func main() {
	// Map operators (runes) to the functions above
	funcMapper := map[rune]func(int, int) int{
		'+': add,
		'-': sub,
		'*': mul,
		'/': div,
	}
	scanner := bufio.NewScanner(os.Stdin)
	var op rune
	var num1, num2 int

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, " ", "", -1) // remove spaces
		if n, _ := fmt.Sscanf(line, "%d%c%d", &num1, &op, &num2); n == 3 {
			if f, present := funcMapper[op]; present {
				fmt.Println(f(num1, num2))
			} else {
				fmt.Println("bad operator:", op)
			}
		} else {
			fmt.Println("bad line:", line)
		}
	}
}
