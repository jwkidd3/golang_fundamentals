// Add a second server goroutine, and use a select statement to send
// the work to whichever server is able to accept it (to test this,
// send a big chunk of work, i.e., a large int, to one of them, and
// then be sure the other can do subsequent work in the meantime)

package main

import (
	"fmt"
	"strconv"
	"time"
)

func server(c chan int, num int) {
	for {
		// Get "work" from main()
		work := <-c
		fmt.Printf("Database server %d working for %d seconds...\n", num, work)
		time.Sleep(time.Duration(work) * time.Second)
		fmt.Printf("Database server %d done\n", num)
	}
}

func main() {
	var work string
	// Create an array of channels to communicate with servers
	var workchan = [2]chan int{
		make(chan int),
		make(chan int),
	}

	// Start "servers"
	go server(workchan[0], 0)
	go server(workchan[1], 1)

	for {
		fmt.Print("Work? ")
		// if they just hit return, ask again
		if n, _ := fmt.Scanln(&work); n == 0 {
			continue
		}
		if work == "quit" {
			fmt.Println("Server quitting")
			break
		}
		if secs, err := strconv.Atoi(work); err == nil {
			select {
			case workchan[0] <- secs:
				fmt.Println("work sent to server 0")
			case workchan[1] <- secs: // ditto
				fmt.Println("work sent to server 1")
			}
		}
	}
}
