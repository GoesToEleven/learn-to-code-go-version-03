package main

import "fmt"

func main() {

	f := outer()
	fmt.Println(f())
}

func outer() func() int {
	return func() int {
		return 42
	}
}
