package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 23
		// If you don't close c, then the second retrieve below will get a deadlock
		close(c)
	}()

	v, ok := <-c
	fmt.Println("Value from channel", v, "status:", ok)

	v, ok = <-c
	fmt.Println("Value from channel", v, "status:", ok)
}
