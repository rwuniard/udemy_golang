package main

import "fmt"

// Declare variable outside a function
// You need to use var keyword
// It will automatically set to 0 if you don't assign the value.
var z int
var a = 3

func main() {
	fmt.Println("z = ", z)
	x := 43
	y := x + 4
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println("Hello World")
	MyHello()
}

func MyHello() {
	fmt.Println("Inside MyHello")
}
