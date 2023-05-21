package main

import "fmt"

type Person struct {
	First string
	Last  string
}

type SecretAgent struct {
	Person
	ltk int
}

func main() {
	p1 := Person{
		First: "Jenny",
		Last:  "Moneypenny",
	}
	fmt.Printf("p1\t%#v\n", p1)

	p2 := SecretAgent{
		Person: Person{"James", "Bond"},
		ltk:    7,
	}
	fmt.Printf("p2\t%#v\n", p2)
	fmt.Println(p2.First, p2.Last, p2.ltk)
	fmt.Println(p2.Person.First, p2.Person.Last, p2.ltk)
}
