package main

import (
	"fmt"
)

func main() {
	//SEQUENCE
	fmt.Println("this is the first statement to run")
	fmt.Println("this is the second statement to run")
	x := 45 // this is the third statement to run
	y := 5  // this is the fourth statement to run
	fmt.Printf(" x=%v \n y=%v\n", x, y)

	// CONDITIONAL
	// if statements
	// switch statements

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else {
		fmt.Println("equal to, or greater than, the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else if x == 42 {
		fmt.Println("equal to the meaning of life")
	} else {
		fmt.Println("greater than the meaning of life")
	}

	/*
		"If" statements specify the conditional execution of two branches
		according to the value of a boolean expression. If the expression evaluates
		to true, the "if" branch is executed, otherwise, if present, the "else" branch is executed.
	*/
	// https://go.dev/ref/spec#If_statements

	/*
		Comparison operators
		Comparison operators compare two operands and yield an untyped boolean value.

		==    equal
		!=    not equal
		<     less
		<=    less or equal
		>     greater
		>=    greater or equal
	*/
	// https://go.dev/ref/spec#Comparison_operators
}
