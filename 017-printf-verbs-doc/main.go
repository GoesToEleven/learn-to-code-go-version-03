package main

import (
	"fmt"
)

func main() {
	const name, age = "Kim", 22
	fmt.Printf("%s is %d years old.\n", name, age)

	// It is conventional not to worry about any
	// error returned by Printf.

	// use printf to print the VALUE and TYPE for each of these variables

	a := 42
	b := 42.04
	var c uint8 = 255
	var d int8 = -128
	var e uint = 9223372036854775807

	f := []int{7, 42, 365, 1024}

	g := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	type person struct {
		first string
		age   int
	}

	h := person{
		first: "Todd",
		age:   45,
	}

	fmt.Printf("This VALUE %v is of TYPE %T\n",a,a)
	fmt.Printf("This VALUE %v is of TYPE %T\n",b,b)
	fmt.Printf("This VALUE %v is of TYPE %T\n",c,c)
	fmt.Printf("This VALUE %v is of TYPE %T\n",d,d)
	fmt.Printf("This VALUE %v is of TYPE %T\n",e,e)
	fmt.Printf("This VALUE %v is of TYPE %T\n",f,f)
	fmt.Printf("This VALUE %v is of TYPE %T\n",g,g)
	fmt.Printf("This VALUE %v is of TYPE %T\n",h,h)
}
