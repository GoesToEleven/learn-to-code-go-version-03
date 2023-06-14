package main

import "fmt"

// value semantics
func addOne(v int) int {
	return v + 1
}

// pointer semantics
func addOneP(v *int) {
	*v += 1
}

func main() {

	// value semantics
	a := 1
	fmt.Println(a)         // 1
	fmt.Println(addOne(a)) // 2
	fmt.Println(a)         // 1

	// pointer semantics
	b := 1
	fmt.Println("value semantics")
	fmt.Println(b) // 1
	addOneP(&b)
	fmt.Println(b) // 2

}
