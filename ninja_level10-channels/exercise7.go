package main

import "fmt"

// See the solution from the instructor.
// ● https://play.golang.org/p/R-zqsKS03P
// ● https://play.golang.org/p/quWnlwzs2z
// ● https://go.dev/play/p/WqYnBC_CiKn
func main() {

	var receivers []chan int

	for i := 0; i < 10; i++ {
		receivers = append(receivers, generateMessage())
	}

	for i := 0; i < 10; i++ {
		for v := range receivers[i] {
			fmt.Print("receiver:", i)
			fmt.Println(" value", v)
		}
	}
}

func generateMessage() chan int {
	sender := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			sender <- i
		}
		close(sender)
	}()

	return sender
}
