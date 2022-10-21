package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 23
	}()

	v, ok := <-c
	fmt.Println("Value from channel", v, "status:", ok)
}
