package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)
	wg.Add(2)

	go send(c)
	go receive(c)

	wg.Wait()
}

func send(sender chan<- int) {
	for i := 0; i < 100; i++ {
		sender <- i
	}
	close(sender)
	wg.Done()
}

func receive(r <-chan int) {
	for v := range r {
		fmt.Println("receiving value:", v)
	}
	wg.Done()
}
