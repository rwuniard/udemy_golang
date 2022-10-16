package main

import "fmt"

func main() {
	foo()
	bar("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)

	x, y := mouse("Ian", "Fleeming")
	fmt.Println(x)
	fmt.Println(y)
}

// func syntax
// func (r receiver) identifier(parameters) (return(s)) { ... }

// func foo
func foo() {
	fmt.Println("This is foo")
}

// Everything is Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello ", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `, says "Hello"`)
	b := false
	return a, b
}
