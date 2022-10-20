package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutine:", runtime.NumGoroutine())

	go sayHello(1)
	go sayHello(2)
	fmt.Println("Goroutine:", runtime.NumGoroutine())
	wg.Wait()
}

func sayHello(index int) {
	for i := 0; i < 10; i++ {
		fmt.Println("I am goroutine:", index)
		runtime.Gosched()
	}

	wg.Done()
}
