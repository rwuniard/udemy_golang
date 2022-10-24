// Package mymath provides ACME inc math solutions
package mymath

// Sum adds an unlimited number of values of type int
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

// NOTES: run in the doc/go_doc folder
// go doc -> Print all the documentation inside example1.go

// go doc Sum -> Print all the documentation of Sum func
// go doc mymath -> Print the documentation of mymath package, in this example i got an error, since I may not have the proper folder structure and filenaming perhaps.
// go doc mymath.Sum
// go doc fmt
