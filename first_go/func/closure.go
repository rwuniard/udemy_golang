package main

import "fmt"

func main() {

	a := incremetor()
	b := incremetor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func incremetor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
