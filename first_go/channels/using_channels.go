package main

import "fmt"

func main() {
	c := make(chan int)
	// send
	go foo(c)

	// receive
	// We don't use goroutine here, since we may have to use WaitGroup to make sure both of them are finished.
	bar(c)
}

// Send
func foo(c chan<- int) {
	c <- 42
}

// Receive
func bar(c <-chan int) {
	fmt.Println("Value from channel:", <-c)
}
