package main

import "fmt"

func main() {
	exercise1()
}

func exercise1() {
	r := foo()
	fmt.Println(r)

	r2, s := bar()
	fmt.Println(r2)
	fmt.Println(s)
}

func foo() int {
	return 1
}

func bar() (int, string) {
	x := 1
	s := "hello"
	return x, s
}
