package main

import "fmt"

func main() {
	hello()
}

func hello() {
	defer foo()
	bar()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
