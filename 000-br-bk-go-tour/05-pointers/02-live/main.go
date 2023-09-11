package main

import "fmt"

func main() {
	a := 7
	fmt.Println(a)
	// see the address
	fmt.Println(&a)

	fmt.Println("-----")

	// store a pointer to an int
	// DECLARE a variable with the identifier `b` will store a value of type 'pointer to an int'
	var b *int
	fmt.Println(b)
	
	fmt.Printf("%T \n",b)
	// fix: nil pointer dereference
	// initialize the pointer by assigning the zero value
	// DECLARE and ASSIGN = initialize
	// we are ASSIGNING here using NEW the zero value of an int, and returnging a *int
	b = new(int)
	fmt.Println(b)
	fmt.Println(*b)
	// fmt.Printf("%T \n",b)

	// dereference: change the value stored in a memory address location
	*b = 42
	fmt.Println(b)
	fmt.Printf("%T \n",b)
	fmt.Println(*b)

	// dereference: see the value stored in a memory address location
	fmt.Println(*b)

	s := "original - "
	fmt.Println(s)
	valSem(s)
	fmt.Println(s)

	fmt.Println("-------")
	fmt.Println("address", &s)
	ptrSem(&s)
	fmt.Println(s)
}

func valSem(s1 string) {
	s1 += "change"
	fmt.Println("inside valSem", s1)
}

func ptrSem(s2 *string) {
	*s2 += "change"
	fmt.Println("address", &s2)
	fmt.Println("inside vptrSem", *s2)
}
