package main

import "fmt"

func main() {
	f := 42.42

	// works
	fmt.Println(f / 2)

	// doesn't work
	// n := 2
	// fmt.Println(f/n)

	// works
	// the type of a constant is defined by how it is used
	const c = 2
	fmt.Println(f / c)
}
