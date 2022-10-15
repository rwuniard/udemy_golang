package main

import "fmt"

func main() {

	a := 8
	fmt.Printf("%d\t\t%b\n", a, a)

	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

	// Iota needs to be used inside constant
	const (
		// Now to use iota
		_   = iota             // this will throw away the iota 0
		kb1 = 1 << (iota * 10) // This will shift 10 times
		mb1 = 1 << (iota * 10)
		gb1 = 1 << (iota * 10)
	)

	fmt.Printf("%d\t\t\t%b\n", kb1, kb1)
	fmt.Printf("%d\t\t\t%b\n", mb1, mb1)
	fmt.Printf("%d\t\t%b\n", gb1, gb1)

	// This will shift 0010 -> 1000
	bit := 2
	bit4 := bit << 2

	fmt.Println(bit4)

	bit = 3
	bit4 = bit << 2
	fmt.Println(bit4)
	fmt.Printf("%d\t\t%b\n", bit, bit)
	fmt.Printf("%d\t\t%b\n", bit4, bit4)
}
