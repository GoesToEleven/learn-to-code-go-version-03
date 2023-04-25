package main

import "fmt"

func main() {
	var x uint8 = 255
	var y int8 = 127

	fmt.Printf("%v is of type %T\n",x,x)
	fmt.Printf("%v is of type %T\n",y,y)
}