package main

import (
	"fmt"
)

func main() {
	fmt.Println("203-hands-on-11")
	type ByteSize float64
	x := "my string"
	y := 1
	var z float64 = 2.0
	var f ByteSize = 3
	fmt.Printf("value: %v, and type: %T \n", x , x)
	fmt.Printf("value: %v, and type: %T \n", y , y)
	fmt.Printf("value: %v, and type: %T \n", z , z)
	fmt.Printf("value: %v, and type: %T \n", f , f)
	
}
