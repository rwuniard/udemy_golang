package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	var err error
	bs, origErr := json.Marshal(a)

	if origErr != nil {
		errorMsg := fmt.Sprintf("Error during marshaling inside toJson func: %v", origErr.Error())
		err = errors.New(errorMsg)
		return bs, err
	}

	return bs, nil
}
