package main

import "fmt"

func main() {
	fmt.Println(factorial(50))
	fmt.Println(loopFact(50))
}

func factorial(x uint64) uint64 {
	// fmt.Println(x)
	if x == 1 {
		return 1
	}
	return x * factorial(x-1)
}

func loopFact(x uint64) uint64 {
	var total uint64
	total = 1
	for ; x > 0; x-- {
		total *= x
	}
	return total
}
