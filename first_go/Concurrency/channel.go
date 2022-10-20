package main

import "fmt"

// When you have a function that returns a value,
// you want to wrap it in another function that will become the go routine
// Get the return value and put it in a channel
func doSomething(x int) int {
	return x * 5
}

func main() {
	ch := make(chan int)
	go func() {
		ch <- doSomething(5)
	}()
	fmt.Println(<-ch)
}
