package main

import (
	"fmt"
)

func main() {
	//SEQUENCE
	fmt.Println("this is the first statement to run")
	fmt.Println("this is the second statement to run")
	x := 40 // this is the third statement to run
	y := 5  // this is the fourth statement to run
	fmt.Printf(" x=%v \n y=%v\n", x, y)

	// CONDITIONAL
	// if statements
	// switch statements

	if x < 42 && y < 42 {
		// && requries both to be true to run
		fmt.Println("both are less than the meaning of life")
	}

	if x > 30 || x < 42 {
		// || requires one of them to be true to run
		fmt.Println("x is getting close to the meaning of life")
	}

	if x != 42 && y != 10 {
		fmt.Println("x is not 42 & y is not 10")
	}

	/*
		Logical operators
		Logical operators apply to boolean values
		and yield a result of the same type as the operands.
		The right operand is evaluated conditionally.

		&&    conditional AND    p && q  is  "if p then q else false"
		||    conditional OR     p || q  is  "if p then true else q"
		!     NOT                !p      is  "not p"
	*/
	// https://go.dev/ref/spec#Logical_operators
}
