package main

import "fmt"

// 30
// 30 is a NUMERIC LITERAL is also an UNTYPED CONSTANT

const a = 42     // untyped constant aka "constant of a kind"
const b int = 43 // typed constant

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T \t %v \n", a, a)
	fmt.Printf("%T \t %v \n", b, b)

	const c = 30
	fmt.Printf("%T \t %v \n", c, c)

	fm := floatMe(30)
	fmt.Printf("%T \t %v \n", fm, fm)
}

func floatMe(f float64) float64 {
	return f * 2.5
}
