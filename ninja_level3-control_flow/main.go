package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
	exercise7()
	exercise8()
	exercise9()
	exercise10()
}

func exercise1() {
	fmt.Println("Exercise1")
	for i := 0; i <= 10000; i++ {
		fmt.Println(i)
	}
}

func exercise2() {
	fmt.Println("Exercise2")
	for x := 65; x <= 90; x++ {
		for i := 0; i < 3; i++ {
			fmt.Printf("%#U\n", x)
		}
	}
}

func exercise3() {
	fmt.Println("Exercise3")
	year := 2022
	for year >= 1973 {
		fmt.Println(year)
		year--
	}

}

func exercise4() {
	fmt.Println("Exercise4")
	year := 1999
	for {
		if year > 2022 {
			break
		}
		fmt.Println(year)
		year++
	}
}

func exercise5() {
	fmt.Println("Exercise5")
	x := 10

	for x <= 100 {
		result := x / 4
		remainder := x % 4

		fmt.Printf("%d/4=%d  reminder:%d\n", x, result, remainder)
		x++
	}
}
func exercise6() {
	fmt.Println("Exercise6")
	x := 10
	if x < 20 {
		fmt.Println("x < 20")
	}
}
func exercise7() {
	fmt.Println("Exercise7")
	x := 10
	if x > 20 {
		fmt.Println("x < 20")
	} else if x == 10 {
		fmt.Println("x == 10")
	} else {
		fmt.Println("none of the conditions are met")
	}
}

func exercise8() {
	fmt.Println("Exercise8")
	switch {
	case 4 == 5:
		fmt.Println("4 == 5")
	case 3 > 2:
		fmt.Println("3 > 2")
	case 3 < 2:
		fmt.Println("3 < 2")
	default:
		fmt.Println("default")
	}
}
func exercise9() {
	fmt.Println("Exercise9")
	s := "Skiing"
	switch s {
	case "Ski":
		fmt.Println("Ski")
	case "Snowboarding":
		fmt.Println("Snowboarding")
	case "Skiing":
		fmt.Println("Skiing")
	default:
		fmt.Println("default")
	}
}
func exercise10() {
	fmt.Println("Exercise10")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
