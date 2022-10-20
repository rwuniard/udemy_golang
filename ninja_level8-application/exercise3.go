package main

import (
	"encoding/json"
	"fmt"
	"os"
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
	// fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
