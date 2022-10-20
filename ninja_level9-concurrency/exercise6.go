package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Go run exercise6.go
	// Go build exercise6.go will build the executable in the current folder
	// Go install exercise6.go -> will put it in the ~/go/bin folder. This is defined in GOROOT env variable.
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)
}
