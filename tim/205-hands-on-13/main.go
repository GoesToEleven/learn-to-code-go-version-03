package main

import (
	"fmt"
)

func main() {
	fmt.Println("205-hands-on-13")
	var a int8 = -128
	var b int8 = 4 //-32768
	var c uint8
	var d int16 = 0
	fmt.Printf("%d \t \t %b\n", a, a)
	fmt.Printf("%d \t \t %b\n", ^a, ^a)
	fmt.Printf("%d \t \t %b\n", b, b)
	fmt.Printf("%d \t \t %b\n", ^b, ^b)
	fmt.Printf("%d \t \t %b\n", c, c)
	fmt.Printf("%d \t \t %b\n", ^c, ^c)
	fmt.Printf("%d \t \t %b\n", d, d)
	fmt.Printf("%d \t \t %b\n", ^d, ^d)

	fmt.Println("| operator")
	var a1 uint8 = 124
	fmt.Printf("%08b\n", a1)
	var b1 uint8 = 196
	fmt.Printf("%08b\n", b1)
	a1 |= b1
	fmt.Printf("%08b\n", a1)

	fmt.Println("Testing")
	fmt.Printf("%08b\n", -128^1)
}
