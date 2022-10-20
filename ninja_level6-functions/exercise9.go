package main

import "fmt"

func main() {

	fromUS(helloFrom)
	fromIndo(helloFrom)

}

func fromUS(hello func(s string) string) int {
	a := 1
	fmt.Println(hello("USA"))

	return a
}

func fromIndo(hello func(s string) string) int {
	a := 62
	fmt.Println(hello("Indonesia"))
	return a
}

func helloFrom(s string) string {
	x := "hello world from " + s
	return x
}
