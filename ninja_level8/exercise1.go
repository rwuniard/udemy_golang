package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName  string
	FirstName string
	LastName  string
	Age       int
}

func main() {
	u1 := User{
		UserName:  "007",
		FirstName: "James",
		LastName:  "Bond",
		Age:       34,
	}
	u2 := User{
		UserName:  "0011",
		FirstName: "Miss",
		LastName:  "Moneypenny",
		Age:       29,
	}

	users := []User{u1, u2}
	results, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(results))

	// Use indentation
	result2, err2 := json.MarshalIndent(users, "", "  ")
	if err2 != nil {
		fmt.Println("Error:", err2)
	}
	fmt.Println(string(result2))
}
