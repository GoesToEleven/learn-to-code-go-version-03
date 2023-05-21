package interfaceimplementation

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

type secretagent struct {
	person
	Number string
}

func (p person) speak() {
	fmt.Println("I'm a civilian and my name is", p.Name)
}

func (sa secretagent) speak() {
	fmt.Println("I work for the Majesty's secret service and my name is", sa.Name, "and my number is", sa.Number)
}

type human interface {
	speak()
}

func saysomething(h human) {
	h.speak()
}

func ThisWorks() {
	// defined type T
	// receiver type T

	p1 := person{"Miss Moneypenny"}
	p1.speak()
	saysomething(p1)

	sa1 := secretagent{person{"James "}, "007"}
	sa1.speak()
	saysomething(sa1)
}

func (p *person) joke() {
	fmt.Println("Here is a joke from a civilian", p.Name)
}

func (sa *secretagent) joke() {
	fmt.Println("Here is a joke from a secret agent", sa.Name, sa.Number)
}

type comedian interface {
	joke()
}

func tellajoke(c comedian) {
	c.joke()
}

func ThisDoesNotWork() {
	// defined type T
	// receiver type *T

	p2 := person{"Miss Moneypenny"}
	// calling the method works
	fmt.Println("Calling the method works")
	p2.joke()
	// using the interface does not work
	fmt.Println("USING THE INTERFACE DOES NOT WORK")
	// tellajoke(p2)

	sa2 := secretagent{person{"James Bond"}, "007"}
	// calling the method works
	fmt.Println("Calling the method works")
	sa2.joke()
	// using the interface does not work
	fmt.Println("USING THE INTERFACE DOES NOT WORK")
	// tellajoke(sa2)
}

func ThisAlsoAllWorks() {
	// defined type T
	// receiver type T

	p3 := person{"Miss Moneypenny"}
	p3.speak()
	saysomething(p3)

	sa3 := secretagent{person{"James Bond"}, "007"}
	sa3.speak()
	saysomething(sa3)

	// defined type *T
	// receiver type T

	p4 := &person{"Miss Moneypenny"}
	p4.speak()
	saysomething(p4)

	sa4 := &secretagent{person{"James Bond"}, "007"}
	sa4.speak()
	saysomething(sa4)

	// defined type *T
	// receiver type *T

	p5 := &person{"Miss Moneypenny"}
	p5.joke()
	tellajoke(p5)

	sa5 := &secretagent{person{"James Bond"}, "007"}
	sa5.joke()
	tellajoke(sa5)
}
