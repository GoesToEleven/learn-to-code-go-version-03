package interfaceimp

import "fmt"

/*
The method set of a defined TYPE T
methods declared with receiver type T.

The method set of a POINTER to a defined TYPE T
(where T is neither a pointer nor an interface)
methods declared with receiver *T or T.
*/

type person struct {
	Name string
}

func (p person) speak() {
	fmt.Println("I'm a civilian and my name is", p.Name)
}

func (p *person) joke() {
	fmt.Println("Here is a joke", p.Name)
}

type human interface {
	speak()
}

type comedian interface {
	joke()
}

func saysomething(h human) {
	h.speak()
}

func tellajoke(c comedian) {
	c.joke()
}

func Runnit() {
	// defined type T, receiver type T

	// this works
	p := person{"James"}
	p.speak()
	p.joke()
	saysomething(p)

	// this doesn't work
	// tellajoke(p)

	pp := &p

	// this works
	pp.speak()
	saysomething(pp)

	// this works
	tellajoke(pp)
}
