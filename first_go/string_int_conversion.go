package main

import (
	"fmt"
	"strconv"
)

func main() {

	n := 39

	// Converting int to string
	chars := strconv.Itoa(n)
	fmt.Println(chars)
	// Chars[0] contains the ascii value of the first digit, then you convert it to string
	n1 := string(chars[0])
	n2 := string(chars[1])

	// Converting string to int
	a1, err := strconv.ParseInt(n1, 10, 64)
	fmt.Printf("a1: %d, err: %v\n", a1, err)

	fmt.Println(n1, n2)
	// converting string to int
	v1, err1 := strconv.Atoi(n1)
	fmt.Println(v1, err1)
	v2, err2 := strconv.Atoi(n2)
	fmt.Println(v2, err2)

	result := v1 + v2
	fmt.Println(result)
}
