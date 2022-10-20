package main

import "fmt"

func main() {
	x := myFunc()
	result := x()
	fmt.Println(result)
}

func myFunc() func() int {
	a := 1
	return func() int {
		fmt.Println("Inside the func:", a)
		return a
	}
}
