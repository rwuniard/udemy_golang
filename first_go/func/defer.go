package main

import "fmt"

func main() {
	// defer foo()
	// bar()
	demoDeferFunc()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func bar2() {
	fmt.Println("bar2")
}

func demoDeferFunc() {
	// foo won't be executed until the demoDeferFunc about to exit
	defer foo()
	bar()
	bar2()
}
