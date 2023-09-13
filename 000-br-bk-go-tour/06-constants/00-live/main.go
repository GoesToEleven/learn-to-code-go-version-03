package main

import "fmt"

// 42
// 42 and all other numbers are NUMERIC LITERALS -- UNTYPED CONSTANTS

const a = 42     // untyped "constant of a kind"
const b int = 42 // typed

func main() {
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("------------")
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)

	fmt.Println("------------")
	fm := floatMe(a)
	fmt.Println(fm)
	fmt.Printf("%T \n", fm)

	fmt.Println("------------ DIVISION")
	c := 4 / 2
	fmt.Printf("%T \t %v \n", c, c)
	d := 4 / 2.1
	fmt.Printf("%T \t %v \n", d, d)
}

func floatMe(f float64) float64 {
	return f * 2.5
}
