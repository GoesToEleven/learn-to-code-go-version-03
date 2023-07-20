package main

import "fmt"

func Add(x, y int) int {
	return x + y
}

func main() {
	sum := Add(42, 13)
	fmt.Println(sum) // Output: 55
}
