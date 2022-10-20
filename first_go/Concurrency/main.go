package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	// This will run before main function.
}

var wg sync.WaitGroup

func main() {
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("ARCH", runtime.GOARCH)
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	// I am not sure what this Add does to the wait group.
	wg.Add(2)
	//This launches the foo into another go routine
	go foo()
	go bar()
	fmt.Println("Goroutines", runtime.NumGoroutine())
	// This will wait until all of them done.
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	// This tells that this guy is done at this point.
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
	wg.Done()
}
