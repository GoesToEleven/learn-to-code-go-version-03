package main

import "fmt"

// person ...
type person struct  {
	f string
}

// new ...
func new(s string) *person {
	return &person{
		f: s,
	}
}

func (p *person) newName() {
	p.f = "new"
}



func main() {
	p1 := new("james")
	p2 := new("jenny")
	fmt.Println(p1, p2)
	p1.newName()
	p2.newName()
	fmt.Println(p1, p2)
}
