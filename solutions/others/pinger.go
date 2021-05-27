package main

import "fmt"
import "time"

func pinger(c chan string) {
	fmt.Println("pinger about to send")
	c <- "ping...ping...ping" // send a value into a channel ... blocking operation
	fmt.Println("Susmitha's print goes here")
}

func main() {
	mychan := make(chan string) // Create a new channel
	go pinger(mychan)           // run pinger() as a goroutine
	fmt.Println("goroutine started...waiting 3 seconds before receiving")
	time.Sleep(3 * time.Second)
	msg := <-mychan // receive value from channel ... this is a blocking operation
	fmt.Println(msg)
	//time.Sleep(1 * time.Second)
}
