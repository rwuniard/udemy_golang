package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 3, 8, 20, 15, 44, 33, 88, 45, 39, 94}
	xs := []string{"James", "Q", "M", "Monty", "Dr. No", "Moneypenny"}
	xf := []float64{3.22, 1.223, 4.55, 8.2, 3.4, 5.4}

	fmt.Println("Before sorting:", xi)
	sort.Ints(xi)
	fmt.Println("After sorting:", xi)

	fmt.Println(xs)

	fmt.Println("Before sorting:", xs)
	sort.Strings(xs)
	fmt.Println("After sorting:", xs)

	fmt.Println("Before sorting:", xf)
	sort.Float64s(xf)
	fmt.Println("After sorting:", xf)
}
