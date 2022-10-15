package main

import "fmt"

func main() {
	myloop()
	nestedLoop()
	forCondition()
	forWithoutCondition()
	breakContinue()
	printAscii()
}

func myloop() {
	fmt.Println("myloop")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func nestedLoop() {
	fmt.Println("nestedLoop")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		for i := 0; i < 3; i++ {
			fmt.Println("inner loop: ", i)
		}
	}
}

func forCondition() {
	fmt.Println("forCondition")
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
}

func forWithoutCondition() {
	fmt.Println("forWithoutCondition")
	x := 1
	for {
		if x > 3 {
			break
		}
		fmt.Println(x)
		x++
	}
}

func breakContinue() {
	fmt.Println("breakContinue")
	x := 0
	for {
		x++
		if x%2 != 0 {
			continue
		} else if x > 10 {
			break
		}
		fmt.Println(x)

	}
}

func printAscii() {
	fmt.Println("printingAscii")
	for i := 33; i <= 122; i++ {
		fmt.Printf("%d\t\tunicode=%#U\n", i, i)
		fmt.Printf("%d\t\tascii=%c\n", i, i)
	}
}
