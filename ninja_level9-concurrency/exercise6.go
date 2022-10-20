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

	atomic.AddInt32(&counter, 1)

	fmt.Println("goroutine:", i, "counter:", atomic.LoadInt32(&counter))

	wg.Done()
}
