package main

import "fmt"

func main() {
	xi := make([]int, 0, 7)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println(xi)
	// copy(xi, []int{1,2,3,4,5,6,7,8}) // values aren't copied into a slice with len zero
	xi = append(xi, 9, 10) // you can append into a slice with len zero or len N
	fmt.Println(xi)
}
