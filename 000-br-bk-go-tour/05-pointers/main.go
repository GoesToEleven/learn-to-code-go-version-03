package main

import "fmt"

// type user struct {
// 	first string
// }

func main() {
	// a := 42
	// initialize = declare + assign

	// fmt.Println(a)
	// fmt.Printf("%T\n", a)

	// show me the address, aka, reference the address
	// fmt.Println(&a)

	// show me the value stored at an address, aka, DEREFERENCE the address
	// fmt.Println(*&a)
	// fmt.Printf("%T \n", a)
	// fmt.Printf("%T \n", &a)

	// *&a = 72
	// fmt.Println(a)

	// DECLARE there is a variable with the identifier `b`
	// to store a VALUE of TYPE int
	// var b int
	// b = a
	// fmt.Println(b)

	// DECLARE ther is a variable with the identifier `c`
	// that stores a VALUE of TYPE *int
	// c := &a
	// fmt.Printf("%T \n", c)
	// b = c

	// d := 7
	// fmt.Println(d)
	// valS(d)
	// fmt.Println(d)

	// ptrS(&d)
	// fmt.Println(d)

	// u1 := ffVS()
	// fmt.Printf("%p \n", &u1)
	// fmt.Println(u1)

	// fmt.Println("----------")

	// u2 := ffPS()
	// fmt.Printf("%p \n", u2)
	// fmt.Println(*u2)
	// fmt.Printf("%T \n", u2)

	// we have DECLARED that there is a variable with the IDENTIFIER of 'e'
	// that will store a VALUE of TYPE 'pointer to an int'
	fmt.Println("-------------")
	var e *int
	fmt.Println(e)
	fmt.Printf("%T \n", e)
	fmt.Printf("%p \n", e)
	fmt.Println("-------------")

	// we have DECLARED e to store a VALUE of TYPE pointer to an int
	// but now we need to ASSIGN a value to that variable in order to
	// have the compiler create memory for us
	// declare + assign = initialize
	e = new(int)
	fmt.Println(e)
	fmt.Println(*e)
	fmt.Printf("%T \n", e)
	fmt.Printf("%p \n", e)
	fmt.Println("-------------")

	// Using the DEREFERENCE operator '*'
	// DEREFERENCE the address stored at 'e' to get to the VALUE
	// and then store a new VALUE at that memory location of 1024
	*e = 1024
	fmt.Println(e)
	fmt.Println(*e)
	fmt.Printf("%T \n", e)
	fmt.Printf("%p \n", e)
	fmt.Println("-------------")

}

// POINTER semantics
// go:noinline
// func ffPS() *user {
// 	u := user{
// 		first: "James",
// 	}
// 	fmt.Printf("%p \n", &u)
// 	return &u
// }

// TODO: add in code review check - return a pointer

// VALUE semantics
// go:noinline
// func ffVS() user {
// 	u := user{
// 		first: "Jenny",
// 	}
// 	fmt.Printf("%p \n", &u)
// 	return u
// }

// func valS(d1 int) {
// 	d1++
// 	fmt.Println("inside valS", d1)
// }

// func ptrS(d2 *int) {
// 	*d2++
// 	fmt.Println("inside ptrS", *d2)
// }
