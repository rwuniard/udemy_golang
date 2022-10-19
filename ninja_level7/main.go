package main

import "fmt"

func main() {
	exercise1()
	exercise2()
}

func exercise1() {
	a := 42

	fmt.Println("address of a:", &a)
	fmt.Println("value of a:", a)
}

type person struct {
	first string
	last  string
}

func exercise2() {

	p1 := person{
		first: "James",
		last:  "Bond",
	}
	fmt.Println("Before:", p1.first)
	fmt.Println("Before:", p1.last)
	changeMe(&p1)

	fmt.Println("After:", p1.first)
	fmt.Println("After:", p1.last)

}

func changeMe(p *person) {
	(*p).first = "Q"
}
