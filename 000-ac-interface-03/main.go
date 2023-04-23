package main

import (
	"fmt"
)

func main() {
	Flexible(42)
	Flexible("James Bond")

	orLikeThis(43)
	orLikeThis("Miss Moneypenny")
}

// empty interface
func Flexible(x interface{}) {
	fmt.Printf("%v is of type %T\n", x, x)
}

// empty interface
func orLikeThis(x any) {
	fmt.Printf("%v is of type %T\n", x, x)
}
// https://pkg.go.dev/builtin#any
