package main

import "fmt"

// Add imports.

// Add user type and provide comment.
// Declare a struct type to maintain information about a user (name, email and age).
type user struct {
	name string
	email string
	age int
}

func main() {

	// Declare variable of type user and init using a struct literal.
	u := user{
		name: "Jenny",
		email: "008@uk.gov",
		age: 27,
	}

	// Display the field values.
	fmt.Println(u)
	fmt.Println(u.name)
	fmt.Println(u.email)
	fmt.Println(u.age)

	// Declare a variable using an anonymous struct.
	anon := struct{
		name string
		email string
		age int
	}{
		name: "James",
		email: "007@uk.gov",
		age: 32,
	}

	// Display the field values.
	fmt.Println(anon)
	fmt.Println(anon.name)
	fmt.Println(anon.email)
	fmt.Println(anon.age)
}
