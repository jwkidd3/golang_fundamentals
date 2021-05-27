package main

import "fmt"
import "time"

var quit = make(chan bool)

func fibonacci(c chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("fib told to quit!")
			return
		}
	}
}

func main() {
	command := ""
	data := make(chan int)

	go fibonacci(data) // run fibonacci up to capacity of channel

	for {
		num := <-data
		fmt.Println(num)
		fmt.Scanf("%s", &command)
		if command == "quit" {
			quit <- true
			break
		}
	}
	time.Sleep(1 * time.Second)
}
