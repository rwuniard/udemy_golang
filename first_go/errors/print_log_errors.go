package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Println("error happened", err)
		// log.Println("error happened", err)
		// log.Fatalln(err) // Log Fatal will call os.Exit(1). Defer function won't be executed.
		panic(err) // Panic stops the current goroutine. Deferred function will run.
		// log.Panic(err)
	}
	log.Println("hello")
}

func foo() {
	fmt.Println("This is a deferred function")
}
