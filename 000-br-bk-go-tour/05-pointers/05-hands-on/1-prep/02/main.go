// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

import "fmt"

type user struct {
	first string
}

func funcName(u *user, s string) {
	u.first = s
}

func main() {

	// Create a variable of type user and initialize each field.
	u1 := user{
		first: "Joe",
	}

	// Display the value of the variable.
	fmt.Println(u1.first)

	// Share the variable with the function you declared above.
	funcName(&u1, "James")

	// Display the value of the variable.
	fmt.Println(u1.first)
}
