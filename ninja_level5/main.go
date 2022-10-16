package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
}

func exercise1() {
	fmt.Println("Exercise1")
	type person struct {
		first_name           string
		last_name            string
		fav_ice_cream_flavor []string
	}

	p1 := person{
		first_name:           "James",
		last_name:            "Bond",
		fav_ice_cream_flavor: []string{"vanilla", "durian", "mango"},
	}

	p2 := person{
		first_name:           "Q",
		last_name:            "Q",
		fav_ice_cream_flavor: []string{"chocolate", "martini", "cherry"},
	}

	fmt.Println(p1.first_name, p1.last_name)
	for i, v := range p1.fav_ice_cream_flavor {
		fmt.Println(i, v)
	}
	fmt.Println(p2.first_name, p2.last_name)
	for i, v := range p2.fav_ice_cream_flavor {
		fmt.Println(i, v)
	}
}

func exercise2() {
	fmt.Println("Exercise2")
	type person struct {
		first_name           string
		last_name            string
		fav_ice_cream_flavor []string
	}

	p1 := person{
		first_name:           "James",
		last_name:            "Bond",
		fav_ice_cream_flavor: []string{"vanilla", "durian", "mango"},
	}

	p2 := person{
		first_name:           "Q",
		last_name:            "Q",
		fav_ice_cream_flavor: []string{"chocolate", "martini", "cherry"},
	}

	m := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}

	if v, ok := m["Q"]; ok {
		fmt.Println(v.first_name, v.last_name, v.fav_ice_cream_flavor)
	}
}

func exercise3() {
	fmt.Println("Exercise3")

	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "blue",
		},
		fourWheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: false,
	}

	fmt.Println(t1)
	fmt.Println(s1)
}

func exercise4() {
	fmt.Println("Exercise4")

	a := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James Smith",
		friends: map[string]int{
			"Bond": 555,
			"Q":    777,
			"M":    888,
		},
		favDrinks: []string{
			"martini",
			"vodka",
			"gin",
		},
	}

	fmt.Println(a)
}
