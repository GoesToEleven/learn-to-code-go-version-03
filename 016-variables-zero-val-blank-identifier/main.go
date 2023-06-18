package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Printf("The type of a is %T and with value %d \n", a, a)

	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f)

	// this would not work
	/*
		b, c, d, e := 0, 1, 2, 3
		fmt.Println(b, c, d)
	*/

	// this does work
	var g int
	fmt.Println(`Value of g is below`)
	fmt.Println(g)
	fmt.Println(`End of value of g`)
	g = 43
	fmt.Println(g)

	// DECLARE a variable to hold a VALUE of a certain TYPE
	// then ASSIGN a VALUE of that TYPE to the variable
	// initializing a variable
	var h int = 44
	fmt.Println(h)
}
