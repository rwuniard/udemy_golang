package main

import "fmt"

type customErr struct {
	customMsg string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.customMsg)
}

func main() {
	myError := customErr{
		customMsg: "This is custom error",
	}
	foo(myError)
}

func foo(e error) {
	// Showing conversion
	fmt.Println("foo ran -", e, "\n", e.(customErr).customMsg)
}
