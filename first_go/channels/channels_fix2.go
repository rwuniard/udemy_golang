package main

import "fmt"

func main() {
	// This will allow the channel to buffer a value, so it won't block.
	// It is not recommended to use buffer channel
	c := make(chan int, 1)
	// This will not block anymore since the channel has a buffer of 1.
	c <- 42

	// But if you put another value, it will block
	// c <- 3

	fmt.Println(<-c)
}
