package main

import "fmt"

// Compute Fibonacci sequence up to n and send results to the channel
func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // close channel to signal all done
}

func main() {
	c := make(chan int, 10) // buffered channel, 10 slots
	go fibonacci(cap(c), c) // run fibonacci up to capacity of channel
	for i := range c {
		fmt.Println(i)
	}
	x := <-c // read from channel c
	fmt.Println(x)
}
