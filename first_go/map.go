package main

import "fmt"

func main() {
	mapAdd()
	mapRange()
	rangeForSlice()
	mapDelete()

}

func mapAdd() {
	fmt.Println("Map Add")
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	fmt.Println(m["James"])
	fmt.Println(m["M"]) // This will print 0

	// TO check whether it found the key in the map, you can do this
	v, ok := m["M"]
	fmt.Println(v, ok)

	if v, ok := m["James"]; ok {
		fmt.Println("my age is ", v)
	}

	// Adding to the map
	m["todd"] = 33
	fmt.Println(m)
}

func mapRange() {
	fmt.Println("Map Range")
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func rangeForSlice() {
	fmt.Println("Range for Slice")
	xi := []int{3, 4, 2, 5, 6, 7}

	for i, v := range xi {
		fmt.Println(i, v)
	}
}

func mapDelete() {
	fmt.Println("Map Delete")
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
		"Todd":            33,
	}
	fmt.Println(m)

	// To delete
	delete(m, "James")
	fmt.Println(m)

	// This doesn't give you any error even though Q doesn't exist.
	delete(m, "Q")
	fmt.Println(m)

	// To check you can use ,ok
	if v, ok := m["Q"]; ok {
		fmt.Println("this should not be printed, because ", v, "doesn't exist")
	}
}
