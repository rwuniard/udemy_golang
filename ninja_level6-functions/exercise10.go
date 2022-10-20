package main

import "fmt"

func main() {
	a := counter()
	fmt.Println(a())
	fmt.Println(a())

	b := counter()
	fmt.Println(b())
	fmt.Println(b())
}

func counter() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
