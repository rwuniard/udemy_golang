package main

import (
	"fmt"
	"math"
)

func main() {
	mySquare := square{
		length: 4,
	}
	fmt.Println("Square area: ", mySquare.area())
	myCircle := circle{
		radius: 4,
	}
	fmt.Println("Circle area: ", myCircle.area())

	info(myCircle)
	info(mySquare)
	fmt.Println("sqrt:", math.Pow(3, 2))
}

type square struct {
	length int
}

type circle struct {
	radius float64
}

func (sq square) area() float64 {
	return float64(sq.length * sq.length)
}

func (cir circle) area() float64 {
	return (math.Pi * math.Pow(cir.radius, 2))
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}
