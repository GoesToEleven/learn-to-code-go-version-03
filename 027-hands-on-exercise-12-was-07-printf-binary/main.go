package main

import "fmt"

func main() {
	x, y, z := 747, 911, 90210

	fmt.Printf("%d \t\t %b \t\t %#X\n",x,x,x)
	fmt.Printf("%d \t\t %b \t\t %#X\n",y,y,y)
	fmt.Printf("%d \t\t %b \t %#X\n",z,z,z)
}