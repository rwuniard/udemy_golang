package main

import "fmt"

var y = 911
var z = "Shaken, not stirred"
var s = `this is to get double quote "quote"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(s)
	fmt.Printf("%T\n", s)

	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n%b\n%x\n", y, y, y)
}
