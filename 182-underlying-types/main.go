package main

import "fmt"

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

type myNumbers interface {
	~int | ~float64
}

func addT[T myNumbers](a, b T) T {
	return a + b
}

type myAlias int

func main() {
	var n myAlias = 42

	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.2, 2.2))

	fmt.Println(addT(n, 2))
	fmt.Println(addT(1.2, 2.2))

}
