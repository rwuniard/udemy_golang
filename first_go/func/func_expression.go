package main

import "fmt"

func main() {

	a := func(x int) {
		fmt.Println("my first func expression, ", x)
	}

	a(43)
}
