package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
		// log.Fatalln(err) // Log Fatal will call os.Exit(1). Defer function won't be executed.
		// panic(err)
	}
	defer f2.Close()
	fmt.Println("check out the log.txt file in the folder")
}
