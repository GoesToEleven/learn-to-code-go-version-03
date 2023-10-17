package main

import "fmt"

type people struct {
	first string
}

func main() {
	// zero value IDIOMATIC
	var a struct{}
	fmt.Printf("a: %T \t %v \n", a, a)

	// zero value not idiomatic
	b := people{}
	fmt.Printf("b: %T \t %v \n", b, b)

	// zero value IDIOMATIC
	var c people
	fmt.Printf("c: %T \t %v \n", c, c)
	
	// zero value IDIOMATIC
	var d []int
	fmt.Printf("d: %T \t %v \n", d, d)
	if d == nil {
		fmt.Println("NIL")
	}
	if len(d) == 0 {
		fmt.Println("LEN == 0")
	}

	// empty slice, not zero value
	e := []int{}
	fmt.Printf("e: %T \t %v \n", e, e)
	if e == nil {
		fmt.Println("NIL")
	}
	if len(e) == 0 {
		fmt.Println("LEN == 0")
	}

}
