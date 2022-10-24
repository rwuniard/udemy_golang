package math_test

func Sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}

	return total
}
