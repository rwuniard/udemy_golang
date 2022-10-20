package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int
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
	x := counter
	runtime.Gosched()
	x++
	counter = x
	fmt.Println("goroutine:", i, "counter:", counter)
	wg.Done()
}
