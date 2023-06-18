package main

import "fmt"

func main() {
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<1, 1<<1)
	fmt.Printf("%d \t %b\n", 1<<2, 1<<2)
	fmt.Printf("%d \t %b\n", 1<<3, 1<<3)
	fmt.Printf("%d \t %b\n", 1<<4, 1<<4)
	fmt.Printf("%d \t %b\n", 1<<5, 1<<5)
	fmt.Printf("%d \t %b\n", 1<<6, 1<<6)

	fmt.Print("\n")
	fmt.Printf("%d \t %b\n", 2, 2)
	fmt.Printf("%d \t %b\n", 3, 3)
	fmt.Printf("%d \t %b\n", 4, 4)
	fmt.Printf("%d \t %b\n", 5, 5)
	fmt.Print("\n")
	fmt.Printf("%d \t %b\n", 3<<1, 3<<1) // 1*2^2 + 1*2^1 + 0*2^0=	6		     110
	fmt.Printf("%d \t %b\n", 3<<2, 3<<2) // 1*2^3 + 1*2^2 + 0*2^1 + 0*2^0= 12    1100
	fmt.Printf("%d \t %b\n", 3<<3, 3<<3)
	fmt.Printf("%d \t %b\n", 3<<4, 3<<4)
	fmt.Printf("%d \t %b\n", 3<<5, 3<<5)
	fmt.Printf("%d \t %b\n", 3<<6, 3<<6)
}
