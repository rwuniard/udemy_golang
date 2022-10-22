package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func main() {
	fmt.Printf("%T\n", ErrNorgateMath)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}

}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// return 0, errors.New("norgate math: square root of negative number")
		// Or you can declare as a variable
		return 0, ErrNorgateMath
	}
	return 42, nil
}

// An example to use Errorf
func usingErrorf() error {
	i := 3
	return fmt.Errorf("%d %T\n", i, i)
}
