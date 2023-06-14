package main

import "fmt"

type person struct {
	first string
}

func (p person) run() {
	fmt.Println(p.first, "I'm running!")
}

func (pp *person) walk(s string) {
	pp.first = s
	fmt.Println(pp.first, "I'm walking.")
}

func main() {
	p := person{"James"}
	p.run()
	p.walk("Jim")

	p2 := &person{"Jenny"}
	p2.run()
	p2.walk("Jen")
}
