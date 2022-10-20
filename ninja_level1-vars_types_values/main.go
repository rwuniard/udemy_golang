package main

import (
	"fmt"
)

var x1 int
var y1 string
var z1 bool

type myowntype int

var a myowntype

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// It will print 0, empty string and false.
	// That's the default value
	fmt.Println(x1, y1, z1)
	fmt.Println(x1)
	fmt.Println(y1)
	fmt.Println(z1)

	x1 = 41
	y1 = "hello"
	z1 = true
	s := fmt.Sprintf("%d, %s, %t", x1, y1, z1)
	fmt.Println(s)

	// %v is printing its value regardless the type
	s = fmt.Sprintf("%v, %v, %v", x1, y1, z1)
	fmt.Println(s)

	exercise4()
	exercise5()
}

func exercise4() {
	fmt.Println("Exercise 4")
	// This should print 0
	fmt.Println(a)
	// Print the type of a -> main.myowntype
	fmt.Printf("%T\n", a)
	a = 42
	fmt.Println(a)
}

func exercise5() {
	a = 42
	x1 = int(a)

	fmt.Println("Exercise 5")
	fmt.Println(x1)
	fmt.Printf("%T\n", x1)
}
