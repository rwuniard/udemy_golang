package main

import (
	"fmt"

	"github.com/rwuniard/udemy_golang/ninja_level12-godoc/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Snoopy",
		age:  dog.Years(4),
	}

	fmt.Println(fido.name, "is ", fido.age, "dog years old")
}
