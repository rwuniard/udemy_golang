package main

import "fmt"

func main() {
	fmt.Println("hello")
	s1 := foo()
	fmt.Println(s1)

	s2 := bar()
	fmt.Printf("s2: %T\n", s2)
	fmt.Println(s2())

	// or we can do it like this
	fmt.Println("Direct call:", bar()())
}

func foo() string {
	s := "Hello World"
	return s
}

func bar() func() int {
	return func() int {
		return 2000
	}
}
