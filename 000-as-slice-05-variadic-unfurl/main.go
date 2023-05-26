package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	y := make([]int, 1)
	fmt.Printf("%#v\t%T\n", x, x)
	fmt.Printf("%#v\t%T\n", y, y)
	y = append(y, x...)
	printInts(y...)
	printInts(x...)
}

func printInts(z ...int) {
	fmt.Printf("%#v\t%T\n", z, z)
}
