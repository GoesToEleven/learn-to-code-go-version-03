package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d \t 0000000000%b\n", 1, 1)
	fmt.Printf("%d \t 000000000%b\n", 1<<1, 1<<1)
	fmt.Printf("%d \t 00000000%b\n", 1<<2, 1<<2)
	fmt.Printf("%d \t 0000000%b\n", 1<<3, 1<<3)
	fmt.Printf("%d \t 000000%b\n", 1<<4, 1<<4)
	fmt.Printf("%d \t 00000%b\n", 1<<5, 1<<5)
	fmt.Printf("%d \t 0000%b\n", 1<<6, 1<<6)
	fmt.Printf("%d \t 000%b\n", 1<<7, 1<<7)
	fmt.Printf("%d \t 00%b\n", 1<<8, 1<<8)
	fmt.Printf("%d \t 0%b\n", 1<<9, 1<<9)
	fmt.Printf("%d \t %b\n", 1<<10, 1<<10)
	// fmt.Printf("%d \t\t %b \t\t %v\n", 1<<10, 1<<10, math.Pow(2, 10))
	// fmt.Printf("%d \t %b \t %v\n", 1<<20, 1<<20, math.Pow(2, 20))
}

// shifted 100 places
// 1267650600228229401496703205376
// 18446744073709551615
