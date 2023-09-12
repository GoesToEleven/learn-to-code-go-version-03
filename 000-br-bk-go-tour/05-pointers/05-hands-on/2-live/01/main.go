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
	// is idiomatic
	a := 20
	fmt.Println(a)

	// not idiomatic
	var b int
	b = 20
	fmt.Println(b)

	fmt.Println("-------")
	
	// Display the address of and value of the variable.
	fmt.Println(&a)
	fmt.Printf("%p \n", &a)
	
	fmt.Println("-------")
	
	// Declare a pointer variable of type int. Assign the
	// address of the integer variable above.
	c := &a
	fmt.Println(c)
	fmt.Printf("%v \t %T \n", c, c)
	fmt.Println("------- EXPLORING C")
	fmt.Println(&c)
	fmt.Printf("%v \t %T \n", &c, &c)
	fmt.Println("------- NIL")
	// nil is the zero value of a pointer
	var d *int
	fmt.Println(d)
	fmt.Printf("%p \t %T \n", &d, d)
	d = &a
	fmt.Println(d)
	fmt.Println(*d)
	*d =21
	fmt.Println("a", a)
	fmt.Println("c", *c)
	fmt.Println("d", *d)
	
	
	// Display the address of, value of and the value the pointer
	// points to.
	
	fmt.Println("------- FINAL ONE")
	// value of
	fmt.Println(c)
	
	//address of a
	fmt.Println(&c)
	
	// the value the pointer points to
	fmt.Println(*c)	
}
