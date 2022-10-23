package main

import (
	"testing"
)

// To Execute this test:
// go test exercise5_test.go
func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

func Average(f []float64) float64 {
	total := 0.0

	for _, v := range f {
		total += float64(v)
	}
	return total / float64(len(f))
}
