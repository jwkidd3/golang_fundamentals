package main

import "fmt"

func main() {
	f := 1.23456789
	i := 123456789
	s := fmt.Sprintf("Float value %.3f and int |%15d|", f, i)
	fmt.Println(s)
}
