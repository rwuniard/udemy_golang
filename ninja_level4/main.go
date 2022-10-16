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

	x := [5]int{}
	for i := range x {
		x[i] = i + 1
	}
	fmt.Println(x)

	for i := range x {
		fmt.Println(x[i])
	}

	fmt.Printf("%T\n", x)
}

func exercise2() {
	fmt.Println("Exercise2")
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range x {
		fmt.Println(x[i])
	}
	fmt.Printf("%T\n", x)
}

func exercise3() {
	fmt.Println("Exercise3")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x1 := x[0:5]
	fmt.Println(x1)

	x2 := x[5:]
	fmt.Println(x2)

	x3 := x[2:7]
	fmt.Println(x3)

	x4 := x[1:6]
	fmt.Println(x4)

}

func exercise4() {
	fmt.Println("Exercise4")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

}

func exercise5() {
	fmt.Println("Exercise5")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[6:]...)
	fmt.Println(y)
}

func exercise6() {
	fmt.Println("Exercise6")

	states := make([]string, 50, 50)
	fmt.Println(states)
	fmt.Println("length:", len(states))
	fmt.Println("capacity:", cap(states))

	stateNames := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`,
		` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`,
		` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`,
		` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`,
		` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`,
		` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`,
		` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	fmt.Println(stateNames)

	// The instruction says that I need to use append, but that won't make states as 50 length.
	// It will have 50 blank + 50 states.
	for i := 0; i < len(stateNames); i++ {
		states[i] = stateNames[i]
	}

	fmt.Println(states)
	fmt.Println("length:", len(states))
	fmt.Println("capacity:", cap(states))

}

func exercise7() {
	fmt.Println("Exercise7")
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	data := [][]string{jb, mp}

	fmt.Println(data)
}

func exercise8() {
	fmt.Println("Exercise8")
	fav := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(fav)
}

func exercise9() {
	fmt.Println("Exercise9")
	fav := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	for k, v := range fav {
		fmt.Println("Key: ", k)
		for i := range v {
			fmt.Println(v[i])
		}

	}
}

func exercise10() {
	fmt.Println("Exercise10")
	fav := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	delete(fav, "bond_james")

	for k, v := range fav {
		fmt.Println("Key: ", k)
		for i := range v {
			fmt.Println(v[i])
		}

	}
}
