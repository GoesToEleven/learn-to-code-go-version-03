package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []float64{3, 1, 4, 2}
	// n := []float64{3, 1, 2}
	// n := []float64{}

	fmt.Println(medianOne(n))
	fmt.Println("after medianOne", n)
	fmt.Println(medianTwo(n))
	fmt.Println("after medianTwo", n)
}

func medianOne(x []float64) (float64, error) {
	if len(x) == 0 {
		return 0.0, fmt.Errorf("median of empty slice")
	}

	// allocate a new underlying array
	n := make([]float64, len(x))
	copy(n, x)

	sort.Float64s(n)
	i := len(n) / 2
	if len(n)%2 == 1 {
		return n[i/2], nil
		// when you divide
		// you get the whole number quotient
		// without the remainder modulus
		// https://go.dev/play/p/2r5WrelUEh7
	}
	return (n[i-1] + n[i]) / 2, nil
}

// uses the same underlying array
// everything is pass by value in go
// the value is being passed into thie func
// and then assigned to x
func medianTwo(x []float64) (float64, error) {
	if len(x) == 0 {
		return 0.0, fmt.Errorf("median of empty slice")
	}

	sort.Float64s(x)
	i := len(x) / 2
	if len(x)%2 == 1 {
		return x[i/2], nil
	}
	return (x[i-1] + x[i]) / 2, nil
}
