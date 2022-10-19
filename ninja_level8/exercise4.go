package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{4, 3, 7, 5, 33, 2, 22, 44, 34, 23, 12}
	s := []string{"james", "James", "Moneypenny", "A", "Q", "a", "M"}

	sort.Ints(i)
	fmt.Println(i)

	sort.Strings(s)
	fmt.Println(s)
}
