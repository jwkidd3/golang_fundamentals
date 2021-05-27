package main

import "fmt"
import "math/rand"
import "sync"
import "time"

var wg sync.WaitGroup

func sleeper(num int) {
	defer wg.Done()
	seconds := rand.Int31n(10) + 5 // 5..14 seconds
	fmt.Printf("Sleeper %d sleeping for %d seconds\n", num, seconds)
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Printf("Sleeper %d done\n", num)
}

func main() {
	rand.Seed(time.Now().Unix()) // seed random number generator
	for i := 1; i <= 100000; i++ {
		wg.Add(1)
		go sleeper(i)
	}
	wg.Wait()
}
