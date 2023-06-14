package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking.")
}

func (d *dog) run() {
	d.first = "Rover"
	fmt.Println("My name is", d.first, "and I'm running.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}

func main() {
	d1 := dog{"Henry"}
	d1.walk()
	d1.run()
	// youngRun(d1) // doesn't run
	
	d2 := &dog{"Padget"}
	d2.walk()
	d2.run()
	youngRun(d2)

}

/*

The idea of the method set is integral to how interfaces are implemented and used in Go.

The method set of a type T consists of all methods with receiver type T. 
These methods can be called using variables of type T.

The method set of a type *T consists of all methods with receiver *T or T 
These methods can be called using variables of type *T.
it can call methods of the corresponding non-pointer type as well 

*/
