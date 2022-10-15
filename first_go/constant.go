package main

import "fmt"

// This is a typed constant
const (
	a int = 55
)

// This is untyped constant
const b = "this is a constant"

func main() {
	fmt.Println(a)
	fmt.Printf("%T", a)

	fmt.Println(b)
	fmt.Printf("%T", b)
}
