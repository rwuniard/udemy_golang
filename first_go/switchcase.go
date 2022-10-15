package main

import "fmt"

func main() {
	noFallthrough()
	withFallthrough()
	defaultcase()
}

func noFallthrough() {
	fmt.Println("NoFallthrough")
	x := 2
	switch {
	case false:
		fmt.Println("this should not print")
	case (x == 2):
		fmt.Println(x)
	case (4 == 4):
		fmt.Println("4 == 4")
	}
}

func withFallthrough() {
	fmt.Println("NoFallthrough")
	x := 2
	switch {
	case false:
		fmt.Println("this should not print")
	case (x == 2):
		fmt.Println(x)
		fallthrough // with fallthrough the statement below will be executed
	case (4 == 4):
		fmt.Println("4 == 4")
		fallthrough
	case (5 == 4):
		fmt.Println("not true") // This will be printed since we have fallthrough
		fallthrough
	case (true):
		fmt.Println("should be true")
	}
}

func defaultcase() {
	fmt.Println("Default case")
	x := 1
	switch {
	case false:
		fmt.Println("this should not print")
	case (x == 2):
		fmt.Println(x)
	case (4 != 4):
		fmt.Println("4 == 4")

	case (5 == 4):
		fmt.Println("not true")

	default:
		fmt.Println("this is default")
	}
}
