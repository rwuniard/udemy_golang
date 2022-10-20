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
	s := `[{"UserName":"007","FirstName":"James","LastName":"Bond","Age":34},{"UserName":"0011","FirstName":"Miss","LastName":"Moneypenny","Age":29}]`

	var users []User

	err := json.Unmarshal([]byte(s), &users)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(users)
}
