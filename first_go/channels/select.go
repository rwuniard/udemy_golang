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
		// This select works like range, it will keep pulling values from multiple channels.
		// The sequence of pulling value is based on the order of the value is being put into the channel
		// if I move q <- 0 to the top prior for loop in the send function. It will quit immediately.
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
