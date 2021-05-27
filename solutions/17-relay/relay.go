package main

import "fmt"
import "time"

var done = make(chan bool) // channel to tell main runners are all done

func runner(msg chan int) {
	num := <-msg
	fmt.Printf("Runner %d starting to run...\n", num)
	if num < 4 {
		go runner(msg) // don't do this if I'm last runner
		fmt.Printf("\tRunner %d at starting line, waiting for baton...\n", num+1)
	}
	time.Sleep(5 * time.Second)
	fmt.Printf("Runner %d has finished running.\n", num)
	if num < 4 {
		fmt.Printf("\tRunner %d passing baton to runner %d\n", num, num+1)
		msg <- num + 1 // send next runner # into channel
	} else {
		done <- true
	}
}

func main() {
	baton := make(chan int)
	go runner(baton)
	fmt.Println("Relay race about to begin...")
	baton <- 1 // starter gun!
	<-done     // wait for last runner
	fmt.Printn("Race over!")
}
