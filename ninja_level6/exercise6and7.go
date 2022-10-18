package main

import "fmt"

func main() {
	func() {
		fmt.Println("hello")
	}()

	a := func(s string) {
		fmt.Println(s)
	}
	a("Hello from a")
}
