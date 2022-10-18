package main

import "fmt"

func main() {
	x := 0
	passByValue(x)
	fmt.Println("x after the call to passByValue:", x)

	passByReference(&x)
	fmt.Println("x after the call to passByReference:", x)
}

func passByValue(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

func passByReference(y *int) {
	fmt.Println(*y)
	*y = 43
	fmt.Println(*y)
}
