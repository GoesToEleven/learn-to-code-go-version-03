package main

import "fmt"

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

}

func mapVals(xi []int, mapFunc func(int) int) []int {
	var ii = make([]int, len(xi))
	for i, v := range xi {
		ii[i] = mapFunc(v)
	}
	return ii
}
