// To build: go build mem.go
package main

import "sync"

var a string
var done bool
var mutex = sync.Mutex{}

func setup() {
	a = "hello, world"
	mutex.Lock()
	done = true
	mutex.Unlock()
}

func main() {
    var copyofdone bool
	go setup()
	for {
		mutex.Lock()
		copyofdone = done
		mutex.Unlock()
		if !copyofdone {
			break
		}
	}
	mutex.Lock()
	println(a)
	mutex.Unlock()
}

