package main

import "fmt"

func main() {

	num := []int{1, 2, 3, 4, 5, 6, 7}
	result := foo(num...)
	fmt.Println("Result:", result)

	result2 := bar(num)
	fmt.Println("Result2:", result2)
}

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	return foo(x...)
}
