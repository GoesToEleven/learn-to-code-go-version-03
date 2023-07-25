package main

import "fmt"

type human struct {
	name  string
	email string
}

type secretAgent struct {
	human
	id string
}

type validate interface {
	identify()
}

func (u human) identify() {
	fmt.Printf("HUMAN: \t %s \t %s \n", u.name, u.email)
}

func (sa secretAgent) identify() {
	fmt.Printf("SA: \t %s \t %s \n", sa.name, sa.email)
}

func runMe(v validate) {
	v.identify()
}

func main() {
	sa := secretAgent{
		human: human{
			name:  "James Bond",
			email: "jbond@uk.gov",
		},
		id: "007",
	}
	sa.identify()
	sa.human.identify()
	runMe(sa)
	runMe(sa.human)
}
