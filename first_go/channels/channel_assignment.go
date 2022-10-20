package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	// You cannot assign from a specific type channel to a bi-directional channel
	// c = cr
	// c = cs

	// But you can do it the other way around
	cr = c
	cs = c

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	// Or you can convert it
	// (c) must be enclosed with the parentheses.
	fmt.Printf("%T\n", (<-chan int)(c))
	fmt.Printf("%T\n", (chan<- int)(c))
}
