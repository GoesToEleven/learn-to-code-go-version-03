package main

import "fmt"

func main() {
	foo()
}

// func (r receiver) identifier(p paremeter(s)) (r return(s)) { code }(argument(s))

func foo() {
	fmt.Println("Foo ran")
}
