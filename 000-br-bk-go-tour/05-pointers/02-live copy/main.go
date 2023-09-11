package main

import "fmt"

func main() {
	a := 7
	fmt.Println(a)
	// see the address
	fmt.Println(&a)

	fmt.Println("-----")

	// store a pointer to an int
	var b *int
	fmt.Println(b)
	fmt.Printf("%T \n", b)
	// fix: nil pointer dereference
	// initialize the pointer by assigning the zero value
	b = new(int)
	fmt.Println(b)
	fmt.Printf("%T \n", b)

	// dereference: change the value stored in a memory address location
	*b = 42
	fmt.Println(b)
	fmt.Printf("%T \n", b)

	// dereference: see the value stored in a memory address location
	fmt.Println(*b)

	s := "original - "
	fmt.Println(s)
	valSem(s)
	fmt.Println(s)

	ptrSem(&s)
	fmt.Println(s)
}

func valSem(s string) {
	s += "change"
}

func ptrSem(s *string) {
	*s += "change"
}
