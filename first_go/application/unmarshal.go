package main

import (
	"encoding/json"
	"fmt"
)

// the `json:"First"` -> called tag
type Person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":33},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	byteData := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", byteData)
	fmt.Println(string(byteData))

	var persons []Person
	err := json.Unmarshal(byteData, &persons)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%+v\n", persons)
	fmt.Println(persons)
}
