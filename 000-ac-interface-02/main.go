package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("from person", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("from secret agent", sa.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		first: "moneypenny",
	}

	sa := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}

	saySomething(p)
	saySomething(sa)
}
