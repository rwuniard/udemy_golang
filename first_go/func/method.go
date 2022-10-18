package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// Function syntax
// func (r receiver) identifier(parameters) (return(s)) { ... }
// This receiver will allow any data with type secretAgent to access speak function
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}
func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
		},
		ltk: false,
	}
	fmt.Println(sa1)

	// This is an example of using the receiver
	// This is attaching a method to a struct basically.
	sa1.speak()
	sa2.speak()
}
