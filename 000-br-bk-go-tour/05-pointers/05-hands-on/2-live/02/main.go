package main

import "fmt"

type user struct {
	first string
}

func newName(u *user, s string) {
	u.first = s
}

func main() {

	u1 := user{
		first: "Joe",
	}

	// Display the value of the variable.
	fmt.Println(u1.first)
	
	// Share the variable with the function you declared above.
	newName(&u1, "James")
	
	// Display the value of the variable.
	fmt.Println(u1.first)
}
