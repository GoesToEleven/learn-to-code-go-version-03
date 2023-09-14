package main

import "fmt"

func main() {
	foo()
}

// func, receiver, identifier, [parameter(s)], [return(s)], code block
func foo() {
	fmt.Println("My first func")
}
