package main

import (
	"fmt"
)

func main() {
	shadowExample()
}

func shadowExample() {
	n := 42
	{
		n := 43
		fmt.Println("inner n:", n)
	}
	fmt.Println("outer n:", n)
}
