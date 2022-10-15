package main

import "fmt"

func main() {
	myloop()
	nestedLoop()
	forCondition()
	forWithoutCondition()
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
