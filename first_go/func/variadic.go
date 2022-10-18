package main

import "fmt"

func main() {

	total := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Total ", total)

	// Unfurling according to Todd
	xi := []int{3, 4, 5, 6, 7}
	x := sum(xi...)
	fmt.Println("Result of passing xi:", x)

	// You can also pass 0 param
	result := sum()
	fmt.Println("Result of passing 0 param: ", result)

	// Call sum2
	sum2("james")
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position,", i, "we are now adding, ", v, " to the total which is now", sum)
	}

	fmt.Println("The total is ", sum)
	return sum
}

// The variadic param must be at the end/final parameter in list. The example below won't work.
// func sum2(x ...int, s string) {
// It will work fine like this
func sum2(s string, x ...int) {
	fmt.Println(x)
	fmt.Println(s)
}
