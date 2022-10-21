package main

import (
	"fmt"
)

func main() {
	// This line won't work, since we cannot receive from the channel in the fmt.Println below.
	//cs := make(chan<- int)
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
