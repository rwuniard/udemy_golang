package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- 53
		}
	}()

	v, ok := <-c
	fmt.Println(v, ok)
	// If you don't close c, then you will keep getting ok = true
	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
