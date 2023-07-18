package main

import "fmt"

func main() {

	// in the file before
	// before we leaned up 'type inference'
	// here we are going to specify the type being passed to our generic function

	loveIt[string]("Hello world")
	loveIt[int](7)
	loveIt[bool](true)
	// loveIt(4.2) // doesn't work

}

type myConstraint interface {
	string | int | bool
}

func loveIt[T myConstraint](a T) {
	fmt.Println(a)
}
