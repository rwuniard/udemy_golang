package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	result := sum(numbers...)
	fmt.Println("result:", result)

	// pass sum function as an argument. You don't need to pass the argument of sum, since you are doing it inside the even function.
	result2 := even(sum, numbers...)
	fmt.Println("result2:", result2)

	result3 := odd(sum, numbers...)
	fmt.Println("result3:", result3)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, xj ...int) int {
	var evenNum []int
	for _, v := range xj {
		if v%2 == 0 {
			evenNum = append(evenNum, v)
		}
	}
	return f(evenNum...)
}

func odd(f func(xi ...int) int, xj ...int) int {
	var oddNum []int
	for _, v := range xj {
		if v%1 == 0 {
			oddNum = append(oddNum, v)
		}
	}
	return f(oddNum...)
}
