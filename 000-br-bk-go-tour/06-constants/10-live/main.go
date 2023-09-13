// Multiply two literal constants into a typed variable and display the value.
package main

import "fmt"

// Add imports.

const (
	a          = 42
	b      int = 43
	server     = "myserver"
	port   int = 8080

	c = 50
	d = 2.5
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(server)
	fmt.Println(port)

	// Divide a constant of kind integer and kind floating point and
	// assign the result to a variable.
	e := c / d
	fmt.Println(e)
}
