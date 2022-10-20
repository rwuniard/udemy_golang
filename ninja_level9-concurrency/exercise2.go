package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type animal struct {
	name string
	age  int
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func (p person) speak() {
	fmt.Println(p.firstName, p.lastName, "is", p.age, "old")
}

func (a *animal) speak() {
	fmt.Println(a.name, "is", a.age, "old")
}

func main() {
	p1 := person{
		"James",
		"Bond",
		33,
	}

	a1 := animal{
		"Snoopy",
		4,
	}

	// saySomething will work if we pass the address of p1 or p1,
	// but if func(p *person) speak() {} is implemented, then it cannot take the p1, it needs to pass &p1
	// This is following this rule
	// Receiver        Values
	// ----------------------
	// (t T)           T and *T
	// (t *T)          *T
	saySomething(&p1)
	saySomething(p1)

	saySomething(&a1)
	// This will fail
	// saySomething(a1)
}
