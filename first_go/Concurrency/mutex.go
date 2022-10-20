package main

import (
	"fmt"
	"runtime"
	"sync"
)

// This is an example of racing condition
// To find out if there is a race condition, run it with
// go run -race race_condition.go
func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GoRoutine:", runtime.NumGoroutine())
	// This will set the max num of CPU and return the previous setting
	// fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(4))

	counter := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	// Create a mutex
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			// lock it, so the counter share variable cannot be accessed by others
			// There are several types of Lock. There RLock
			mu.Lock()
			v := counter
			// This allows go to context switch, allowing other goroutines to run.
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		// fmt.Println("GoRoutine:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRoutine:", runtime.NumGoroutine())
	fmt.Println("Count:", counter)
}
