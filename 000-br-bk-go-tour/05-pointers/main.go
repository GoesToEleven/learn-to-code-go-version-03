package main

import "fmt"

type user struct {
	first string
}

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

	u1 := createUserVS()
	fmt.Println(u1)
	fmt.Println(&u1)

	fmt.Println("----------")
	u2 := createUserPS()
	// fmt.Println(*u2)
	fmt.Println(u2)
}

//go:noinline
func createUserPS() *user {
	u := user{
		first: "James",
	}

	println("V2", &u)

	return &u
}

//go:noinline
func createUserVS() user {
	u := user{
		first: "James",
	}

	println("V1", &u)

	return u
}

// func valS(d1 int) {
// 	d1++
// 	fmt.Println("inside valS", d1)
// }

// func ptrS(d2 *int) {
// 	*d2++
// 	fmt.Println("inside ptrS", *d2)
// }
