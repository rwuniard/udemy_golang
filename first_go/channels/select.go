package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// Send
	go send(eve, odd, quit)
	// Receive
	receive(eve, odd, quit)
	fmt.Println("end program")
}

func receive(e, o, q <-chan int) {
	for {
		// This select works like range, it will keep pulling values from multiple channels
		select {
		case v := <-e:
			fmt.Println("from eve channel:", v)
		case v := <-o:
			fmt.Println("from odd channel:", v)
		case v := <-q:
			fmt.Println("from quit channel:", v)
			return
		}
	}
}
func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0
	close(q)
}
