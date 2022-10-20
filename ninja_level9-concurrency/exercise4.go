package main

import (
	"fmt"
	"sync"
)

var counter int
var wg sync.WaitGroup
var mu sync.Mutex

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
	// Lock using mutex
	mu.Lock()
	x := counter
	x++
	counter = x
	fmt.Println("goroutine:", i, "counter:", counter)
	// Unlock it when done.
	mu.Unlock()

	wg.Done()
}
