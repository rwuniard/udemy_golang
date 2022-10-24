package math_test

import (
	"fmt"
	"testing"
)

// To run
// go test -> it will run the math_test.go
// go test -v -> this is the verbose mode.
func TestSum(t *testing.T) {
	x1 := []int{1, 3, 4}
	x2 := []int{3, 5, 3, 4}

	result := Sum(x1...)
	expected_result1 := 8
	if result != expected_result1 {
		t.Error("Expected result:", expected_result1, " but got:", result)
	}

	result = Sum(x2...)
	// This will intentionally to make test to fail
	expected_result2 := 15
	if result != expected_result2 {
		t.Error("Expected result:", expected_result1, " but got:", result)
	}
}

func TestSum2(t *testing.T) {
	// Using a struct to make it easier to compile a test set.
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{21, 21}, 42},
		{[]int{1, 2, 3}, 6},
	}

	for _, test_set := range tests {
		result := Sum(test_set.data...)
		if test_set.answer != result {
			t.Error("Expected result:", result, " but got:", test_set.answer)
		}
	}
}

// This will be showing up in the godoc as an example.
func ExampleSum() {
	x1 := []int{1, 3, 4}

	result := Sum(x1...)
	fmt.Println("The result of Sum:", result)
	// Output:
	// The result of Sum: 8
}

// To run Benchmark
// go test -bench .
// or go test -bench=.
func BenchmarkSum(b *testing.B) {
	x1 := []int{1, 3, 4}
	for i := 0; i < b.N; i++ {
		Sum(x1...)
	}

}
