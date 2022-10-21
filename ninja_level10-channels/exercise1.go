package main

import "fmt"

func main() {
	// Create a channel with buffer 0
	c := make(chan int)

	go func() {
		// This will be caught by the main goroutine, so it wont' blocks
		c <- 42
	}()

	fmt.Println(<-c)
}
