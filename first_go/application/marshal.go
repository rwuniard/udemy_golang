package main

import (
	"encoding/json"
	"fmt"
)

// If the field First is using lower case -> first, it won't print the value.
type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   33,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(string(bs))
}
