package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	defer close(q)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c <-chan int, q chan int) {
	// This will pull all the values from c channel first, then
	// it will pull the value from q.
	// I think this is determine by the sequence of when the value is put into the channel.
	// So, if I put q <-1 at the top of for loop in the gen function, it will quit immediately.
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {

		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}
