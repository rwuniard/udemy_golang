package main

import "fmt"

// Iota will auto increment the number.
const (
	a = iota
	b = iota
	c = iota
	// you can declare it with this too. Just assign iota to the first variable
	// a = iota
	// b
	// c
)

// This will reset the iota number back to 0
const (
	d = iota
	e
	f
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
