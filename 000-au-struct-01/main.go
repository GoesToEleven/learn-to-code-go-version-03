package main

import "fmt"

type Person struct {
	First string
	Last  string
}

func main() {
	var p1 Person
	fmt.Printf("p1\t%#v\n\n", p1)

	p2 := Person{
		First: "James",
		Last:  "Bond",
	}
	fmt.Printf("p2\t%#v\n\n", p2)

	p3 := Person{First: "Q"}
	fmt.Printf("p2\t%#v\n\n", p3)

	p4 := Person{"Miss", "Moneypenny"}
	fmt.Printf("p3\t%#v\n\n", p4)
}
