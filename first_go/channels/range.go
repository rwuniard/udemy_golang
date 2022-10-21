package main

import (
	"fmt"
	"sync"
)

var w sync.WaitGroup

func main() {
	w.Add(2)
	c := make(chan int)

	// send
	go foo(c)

	// receive
	go bar(c)
	w.Wait()
}

// Send
func foo(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	// If you don't close c, the program will have a deadlock
	close(c)
	w.Done()
}

// Receive
func bar(c <-chan int) {
	for channelValue := range c {
		fmt.Println("Value:", channelValue)
	}
	w.Done()
}
