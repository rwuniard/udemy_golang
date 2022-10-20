package main

import "fmt"

func main() {
	c := make(chan int)
	// This will block
	// Because the channel needs to be handed of to another goroutine, but there is no other goroutines, except main
	c <- 42

	fmt.Println(<-c)
}
