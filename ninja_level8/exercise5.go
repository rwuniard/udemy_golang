package main

import (
	"fmt"
	"sort"
)

type User struct {
	UserName  string
	FirstName string
	LastName  string
	Age       int
}

type ByUserName []User

func (b ByUserName) Len() int           { return len(b) }
func (b ByUserName) Less(i, j int) bool { return b[i].UserName < b[j].UserName }
func (b ByUserName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

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
	u3 := User{
		UserName:  "000",
		FirstName: "M",
		LastName:  "Nobody",
		Age:       64,
	}

	users := []User{u1, u2, u3}

	sort.Sort(ByUserName(users))
	fmt.Println(users)
}
