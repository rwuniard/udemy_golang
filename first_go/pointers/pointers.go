package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	//This will print the address of a
	fmt.Println(&a)

	// This will print the type of a
	fmt.Printf("%T\n", a)
	// Print out the type of the &a, should be *int
	fmt.Printf("%T\n", &a)

	// If we assign it to another variable, it will pass the address, so b and a are NOT pointing to the same memory location
	b := a
	fmt.Println("b address:", &b)

	// This is another way to pass the address
	var c *int = &a
	fmt.Println(c)
	// To print the value
	fmt.Println(*c)

	// When you assign another value, since c is pointing to the same address as a, they both will have the same value
	*c = 45
	fmt.Println("c now:", *c)
	fmt.Println("a is now: ", a)

}
