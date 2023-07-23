package main

import "fmt"

func main() {
	xs := []string{"skis", "surf", "snow"}
	fmt.Println(contains(xs, "skis"))
	fmt.Println(contains(xs, "sumo"))

	xi := []int{4, 5, 7}
	fmt.Println(contains(xi, 7))
	fmt.Println(contains(xi, 42))
}

func contains[T comparable](s []T, v T) bool {
	for _, vs := range s {
		if vs == v {
			return true
		}
	}
	return false
}
