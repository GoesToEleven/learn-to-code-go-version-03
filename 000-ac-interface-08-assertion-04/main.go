package main

import "fmt"

func main() {
	// bypasses go's type system
	// rule of thumb: don't do this
	var i any
	i = 42
	fmt.Println(i)
	i = "happiness"
	fmt.Println(i)

	s := i.(string) // type assertion
	fmt.Println(s)

	// this doesn't work
	// can't make assertion of type int for value of type string
	// n := i.(int)
	// fmt.Println(n)

	if n, ok := i.(int); !ok {
		fmt.Printf("can't use int assertion for type %T\n", i)
	} else {
		fmt.Println(n)
	}

	switch i.(type) {
	case int:
		fmt.Println("an int")
	case string:
		fmt.Println("a string")
	default:
		fmt.Printf("some other type: %T\n", i)
	}
}
