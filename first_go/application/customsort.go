package main

import (
	"fmt"
	"sort"
)

// Read this article using sort.Slice and also the sort.Interface
// https://www.callicoder.com/golang-sorting-custom-functions/

type Person struct {
	first string
	age   int
}
type ByAge []Person

// These three functions are defined in the sort.Interface, so if you implement this for any of your data structure,
// you can sort it in any way you want by implementing these functions.
// The sort.Sort(data Interface) -> the Sort will call/use these three functions based on the spec on Interface,
// therefore anything that is implementing these three function can utilize sort.Sort.
func (ba ByAge) Len() int           { return len(ba) }
func (ba ByAge) Swap(i, j int)      { ba[i], ba[j] = ba[j], ba[i] }
func (ba ByAge) Less(i, j int) bool { return ba[i].age < ba[j].age }

// This is to sort it by name
type ByName []Person

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].first < bn[j].first }

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println("After sorting by age:", people)

	sort.Sort(ByName(people))
	fmt.Println("After sorting by name:", people)
}
