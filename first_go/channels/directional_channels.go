package main

import "fmt"

func main() {

	c := make(chan<- int, 2)
	// d := make(<-chan int, 2)
	c <- 42
	c <- 43

	// This will generate compilation error
	// d <- 2
	// d <- 3

	// fmt.Println(<-d)
	// fmt.Println(<-d)
	// This will generate a compiler error
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	fmt.Println("-------")
	fmt.Printf("%T\n", c)
}
