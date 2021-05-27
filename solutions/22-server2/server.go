// - still more...how about a load balancer goroutine which spins up new
//   server routines as needed up to some maximum (i.e., if it sees that
//   no servers are available, it spins up more
// - to do this effectively you'll likely need reflection, so let's
//   simplify by just trying to send to each server, and if they are
//   all busy, launch a new goroutine
// - you'll likely need an array/slice of channels
// - make it so each server goroutine shuts down if it hasn't done work
//   for a while, so even if you hit your max of, say 10, servers, they
//   should start shutting down as the workload decreases
//
// This solution does not use a load balance goroutine (that's left as
// an exercise for you), but it does shut down inactive servers and it
// uses a list of servers (channels) protected by a Mutex.
package main

import (
	"chanlist" // package I wrote to manage of channels
	"fmt"
	"strconv"
	"time"
)

const (
	minServers    = 2
	maxServers    = 10
	serverTimeout = 10 // seconds
)

// a slice of channels to communicate with servers
var workchan chanlist.Chanlist

func server(c chan int, num int) {
	for {
		var work int

		select {
		case work = <-c:
			break
		// if no work received after "serverTimeout" seconds, shut down
		case <-time.After(serverTimeout * time.Second):
			if num >= minServers {
				fmt.Println("Server", num, "idle; shutting down")
				workchan.DeleteChannel(c, minServers)
			}
			return
		}
		fmt.Printf("Database server %d working for %d seconds...\n", num, work)
		time.Sleep(time.Duration(work) * time.Second)
		fmt.Printf("Database server %d done\n", num)
	}
}

// Find a free server to send work to. We acquire the lock on the channel
// list, so we can iterate through the list. Then we do a non-blocking send
// to each channel in turn. If we send successfully, we return the channel
// and true to indicate success. If all channels would block we return an
// empty channel and false indicating failure.
func findFreeServer(work int) (chan int, bool) {
	var emptyChannel chan int
	workchan.Lock() // lock channel list so we can iterate through it

	for index, channel := range workchan.List {
		select {
		case channel <- work: // if we get here, message was sent
			fmt.Printf("found free channel %#v\n", channel)
			return channel, true
		default:
			fmt.Println("Unable to send to", index)
		}
	}
	// channel will be bogus here, but we have to return one
	return emptyChannel, false
}

func main() {
	var ok bool
	var work string
	var index int
	var channel chan int
	workchan.Initialize() // initialize the channel list

	// start "servers"
	go server(workchan.List[0], 0)
	go server(workchan.List[1], 1)

	for {
		fmt.Println(&workchan)
		fmt.Print("> ")
		// if they just hit return, continue
		if count, _ := fmt.Scanln(&work); count == 0 {
			continue
		}
		if work == "quit" {
			fmt.Println("Server quitting")
			break
		}
		if secs, err := strconv.Atoi(work); err == nil {
			// Find a free server, i.e, the channel on which we
			// communicate with it. ok will be false if none available.
			if channel, ok = findFreeServer(secs); !ok {
				// List will still be locked at this point, so call
				// AddChannel to add a channel. (The list will be
				// unlocked when function returns.)
				channel, index = workchan.AddChannel(secs)
				// spin up new "server"
				go server(channel, index)
				// send work to new server
				channel <- secs
			} else {
				workchan.Unlock()
			}
		}
	}
}
