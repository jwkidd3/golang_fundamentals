// A package to contain a "Chanlist" object. This is a slice of
// channels which grows and shrinks as needed. Growing is easy,
// we just use append(), but when shrinking, we likely need to
// remove an item from the middle and "collapse" the slice down.
package chanlist

import (
	"fmt"
	"sync"
)

const initSize = 2

// A slice of channels plus a Mutex to prevent the
// slice from being accessed by multiple goroutines
type Chanlist struct {
	List  []chan int
	Mutex sync.Mutex
}

// Constructor, i.e., create a channel list
func (cl *Chanlist) Initialize() {
	for i := 0; i < initSize; i++ {
		cl.List = append(cl.List, make(chan int))
	}
}

// Print out the channen list nicely
func (cl *Chanlist) String() string {
	var s string
	cl.Lock()
	for _, c := range cl.List {
		s += fmt.Sprintf("%#v\n", c)
	}
	cl.Unlock()
	return s
}

// Delete a channel from the channel list.
func (cl *Chanlist) DeleteChannel(c chan int, min int) {
	cl.Lock()
	fmt.Printf("deleting %#v\n", c)
	num := -1

	// Find the channel in the list
	for index, channel := range cl.List {
		if channel == c {
			num = index
			break
		}
	}
	if num >= 0 {
		// Remove the channel appending the slice up to but not
		// including it, to the slice after it
		cl.List = append(cl.List[:num], cl.List[num+1:]...)
	}
	cl.Unlock()
}

// Add a channel to the list. Note that the list Mutex will
// already be locked when we call this method.
func (cl *Chanlist) AddChannel(work int) (chan int, int) {
	numChans := len(cl.List)
	cl.List = append(cl.List, make(chan int))
	channel := cl.List[numChans]
	fmt.Println("Spinning up server", numChans)
	cl.Unlock()

	return channel, numChans
}

// Just a wrapper function so we can see when the list
// is locked...
func (cl *Chanlist) Lock() {
	fmt.Println("locking chanlist")
	cl.Mutex.Lock()
}

// Just a wrapper function so we can see when the list
// is unlocked...
func (cl *Chanlist) Unlock() {
	fmt.Println("unlocking chanlist")
	cl.Mutex.Unlock()
}
