package main

import "fmt"

// https://pkg.go.dev/builtin@go1.20.3#comparable

func areEqual[T comparable](x, y T) bool {
	return x == y
}

func main() {
	fmt.Println(areEqual(42, 43))
	fmt.Println(areEqual(42, 42))
	fmt.Println(areEqual("Happiness", "Happiness"))

	type secretAgent struct {
		first string
		last  string
	}

	sa1 := secretAgent{"James", "Bond"}
	sa2 := secretAgent{"James", "Bond"}
	fmt.Println(areEqual(sa1, sa2))

	fmt.Println(areEqual(true, true))
	fmt.Println(areEqual(true, false))

	// works
	fmt.Println(areEqual([3]int{1, 2, 3}, [3]int{1, 2, 3}))
	fmt.Println(areEqual([3]int{1, 2, 3}, [3]int{4, 5, 6}))

	// does not work
	// fmt.Println(areEqual([]int{1, 2, 3}, []int{4, 5, 6}))

}
