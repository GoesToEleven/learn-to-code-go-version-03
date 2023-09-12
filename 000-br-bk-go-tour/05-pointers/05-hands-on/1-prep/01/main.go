// Declare and initialize a variable of type int with the value of 20. Display
// the _address of_ and _value of_ the variable.
//
// Declare and initialize a pointer variable of type int that points to the last
// variable you just created. Display the _address of_ , _value of_ and the
// _value that the pointer points to_.
package main

import "fmt"

// Add imports.

func main() {

	// Declare an integer variable with the value of 20.
	a := 20

	// Display the address of and value of the variable.
	fmt.Println(&a)
	fmt.Printf("%p \n", &a)

	// Declare a pointer variable of type int. Assign the
	// address of the integer variable above.
	b := &a
	fmt.Printf("B - %T \t %v \n", b, b)
	
	c := new(int)
	c = &a
	fmt.Printf("C - %T \t %v \n", c, c)
	
	// Display the address of, value of and the value the pointer
	// points to.
	fmt.Println("address: ", &c)
	fmt.Printf("MAS - %T \t %v \n", &c, &c)
	fmt.Println("value of: ", c)
	fmt.Println("value pointer points to: ", *c)
}
