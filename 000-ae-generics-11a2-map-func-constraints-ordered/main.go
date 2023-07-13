package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

/*
In many programming languages,
map is the name of a higher-order function
that applies a given function to each element of a collection,
e.g. a list or set, returning the results in a collection of the same type.
It is often called apply-to-all when considered in functional form.
*/
// https://en.wikipedia.org/wiki/Map_(higher-order_function)

func main() {
	// goal:
	// input: [1,2,3]
	// output: [2,4,6]

	xi1 := []int{1, 2, 3}
	xi2 := mapVals(xi1, func(n int) int {
		return n * 2
	})
	fmt.Println(xi1)
	fmt.Println(xi2)

	// goal:
	// input: [1.4,2.4,3.4]
	// output: [2.8,4.8,6.8]

	xf1 := []float64{1.4, 2.4, 3.4}
	xf2 := mapVals(xf1, func(f float64) float64 {
		return f * 2
	})
	fmt.Println(xf1)
	fmt.Println(xf2)
}

func mapVals[T constraints.Ordered](xi []T, mapFunc func(T) T) []T {
	var ii = make([]T, len(xi))
	for i, v := range xi {
		ii[i] = mapFunc(v)
	}
	return ii
}
