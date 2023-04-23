package main

import "fmt"

func main() {
	do(21)
	do("hello")
	do(true)
}

func do(x any) {
	switch v := x.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}