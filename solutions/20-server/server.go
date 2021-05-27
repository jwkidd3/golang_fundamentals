// Concurrency Exercise:
// - create a program which uses a goroutine to simulate a database server
// - your "database server" should simply be a goroutine which accepts
//   an integer representing some "work" to do and it sleeps for that
//   amount of time before accepting more work
// - your main routine should get integers from the user, and pass those to
//   the "server"
package main

import (
	"fmt"
	"strconv"
	"time"
)

// Run forever and accept "work" on a channel
func server(c chan int) {
	for {
		// Get "work" from main()
		work := <-c
		fmt.Printf("Database server working for %d seconds...\n", work)
		time.Sleep(time.Duration(work) * time.Second)
	}
}

func main() {
	// Create a channel on which to send work. We'll make
	// it a buffered channel so work can stack up a bit
	// before main() blocks.

	workchan := make(chan int, 5)
	work := ""

	// start "server"
	go server(workchan)

	for {
		fmt.Print("Work? ")
		fmt.Scanln(&work) // get "work"

		if work == "quit" {
			fmt.Println("Server quitting")
			break
		}
		// convert to int, and if successful, send the work
		if secs, err := strconv.Atoi(work); err == nil {
			workchan <- secs
		}
	}
}
