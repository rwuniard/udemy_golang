package main

import (
	"fmt"
)

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
}

func exercise1() {
	a := 15

	fmt.Printf("%d\t%b\t%#X\n", a, a, a)
}

func exercise2() {
	b := (42 == 42)
	c := (42 <= 42)
	d := (42 >= 42)
	e := (42 != 42)
	f := (42 < 42)
	g := (42 > 42)

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}

func exercise3() {
	const (
		x int    = 1
		y string = "this is a constant"
	)

	const (
		a = 1
		b = 3.4
	)

	fmt.Println(x, y)
	fmt.Println(a, b)
	fmt.Printf("%T\t%T\n", a, b)
}

func exercise4() {
	a := 42
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)

	b := a << 1
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)
}

func exercise5() {
	a := `This is a string
	literal and "it can be" %d 
	anything`
	fmt.Println(a)
}

func exercise6() {

	const (
		year1 = iota + 2019
		year2
		year3
		year4
	)
	fmt.Println(year1)
	fmt.Println(year2)
	fmt.Println(year3)
	fmt.Println(year4)
}
