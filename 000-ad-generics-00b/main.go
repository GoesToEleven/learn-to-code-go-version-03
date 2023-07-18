package main

import "fmt"

func main() {
	loveIt("Hello world")
	loveIt(7)
	loveIt(true)
	// loveIt(4.2) // doesn't work
}

type myConstraint interface {
	string | int | bool
}

func loveIt[T myConstraint](a T) {
	fmt.Println(a)
}
