package main

import "fmt"

func main() {
	// This will make a channel with buffer of 1
	c := make(chan int, 1)
	// So, when we put a value, it won't block.
	// Since it doesn't have a goroutines to receive it.
	c <- 42

	fmt.Println(<-c)
}
