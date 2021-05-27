package main

import "fmt"
import "time"

func main() {
	// START OMIT
	messages := make(chan string)
	select { // non-blocking send // HL
	case messages <- "hi":
		fmt.Println("sent message")
	default: // HL
		fmt.Println("no message sent")
	}

	go func() {
		messages <- "Here's a message"
	}()
	time.Sleep(1 * time.Second)

	select { // non-blocking receive // HL
	case msg := <-messages:
		fmt.Println("received message:", msg)
	default: // HL
		fmt.Println("no message received")
	}
	// END OMIT
}
