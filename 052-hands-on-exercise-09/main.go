package main

import "fmt"

var x = 40
const y = 41

func main() {
	fmt.Printf("the value of x is %v and the type of x is %T\n",x,x)
	fmt.Printf("the value of y is %v and the type of y is %T\n",y,y)
	z := 42
	fmt.Printf("the value of z is %v and the type of z is %T\n",z,z)
}