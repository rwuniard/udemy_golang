package main

import "fmt"

var a int

// Creating my own type hotdog as an int
type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Printf("%T\n", b)

	// Conversion in Go, other language is casting
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T", a)
}
