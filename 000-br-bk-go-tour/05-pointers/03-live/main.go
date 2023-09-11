package main

import "fmt"

func main() {
	a := 7
	fmt.Println(a)
	fmt.Printf("%T \n", a)
	// show me the address, aka, take the address
	fmt.Println(&a)
	// show me the value at an address, aka, DEREFERENCE the address
	fmt.Println(*&a)
}
