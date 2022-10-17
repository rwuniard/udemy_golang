package main

import "fmt"

func main() {
	foo()

	// The () at the end is to actually call the anonymous func
	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(x int) {
		fmt.Println("The anonymous func with int param is called. Param:", x)
	}(42)
}

func foo() {
	fmt.Println("foo ran")
}
