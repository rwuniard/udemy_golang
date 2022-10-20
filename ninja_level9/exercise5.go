package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int32
var wg sync.WaitGroup

func main() {
	counter = 0
	const numGoRoutine = 50
	wg.Add(numGoRoutine)

	for i := 0; i < numGoRoutine; i++ {
		go incrementor(i)
	}
	wg.Wait()
	fmt.Println("final counter:", counter)
}

func incrementor(i int) {
	// Use atomic with int32
	atomic.AddInt32(&counter, 1)
	// Load the share variable using atomic.LoadInt32
	fmt.Println("goroutine:", i, "counter:", atomic.LoadInt32(&counter))

	wg.Done()
}
