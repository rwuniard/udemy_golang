package main

import "fmt"

func main() {
	x := 43
	y := x + 4
	fmt.Println(x)
	fmt.Println("Hello World")
	MyHello()
}

func MyHello() {
	fmt.Println("Inside MyHello")
}
