package main

import "fmt"

func main() {
	// array literal

	// array zero value

	// array assignment by index position

	// array size is part of type

	// for loop

	// for range clause

	// for range show just the value using the blank identifier (underscore character)

	// for range show just the index

	// CONTIGUOUS MEMORY
	// address comparision: RANGE VALUE vs. INDEX POSITION VALUE

	// SEQUENTIAL CODE
	fmt.Println("-------- psychologists")
	
	// The for range cause seems to be like a function
	// It receives the data structure to range over
	// and then it ranges over it
	
	// range over a VALUE
	psychologists := [5]string{"Alpert", "C.Jung", "Piaget", "Rogers", "Skinner"}

	
	// range over a POINTER
	psychologists = [5]string{"Alpert", "C.Jung", "Piaget", "Rogers", "Skinner"}

}
