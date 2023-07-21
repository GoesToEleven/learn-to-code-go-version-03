package main

import "fmt"

func main() {
	a := make([]int, 0, 10)
	b := make([]int, 10, 10)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("******")
	a = append(a, 0, 1, 2)
	fmt.Println(a)
	b = append(b, 0, 1, 2)
	fmt.Println(b)
	b[3] = 42
	fmt.Println(b)

	// doesn't work
	// a[3] = 42
	// fmt.Println(a)
}
