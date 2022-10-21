package main

import (
	"fmt"
)

func main() {
	// This line won't work, since it initialize receiving channel,
	// therefore you cannot send a value into the channel.
	// cr := make(<-chan int)
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
