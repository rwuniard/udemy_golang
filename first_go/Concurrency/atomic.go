package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// This is an example of racing condition
// To find out if there is a race condition, run it with
// go run -race race_condition.go
func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GoRoutine:", runtime.NumGoroutine())
	// This will set the max num of CPU and return the previous setting
	// fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(4))

	var counter int64 = 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			// We are passing the variable that will be shared and passing the delta.
			// In this case we will be incrementing it by 1.
			atomic.AddInt64(&counter, 1)
			// This allows go to context switch, allowing other goroutines to run.
			runtime.Gosched()
			// To read the variable you need to use LoadInt64
			fmt.Println("Counter:", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		// fmt.Println("GoRoutine:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRoutine:", runtime.NumGoroutine())
	fmt.Println("Count:", counter)
}
