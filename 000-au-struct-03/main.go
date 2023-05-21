package main

import "fmt"

type Person struct {
	First string
	Last  string
}

func (p Person) Speak(s string) {
	fmt.Printf("%s says %s\n", p.First, s)
}

func (p *Person) NewName(s string) {
	p.First = s
}

func main() {
	p1 := Person{
		First: "James",
		Last:  "Bond",
	}
	fmt.Printf("p1\t%#v\n", p1)
	p1.Speak("Shaken, not stirred.")
	p1.NewName("Jim")
	fmt.Printf("%#v\n\n", p1)

	p2 := &Person{
		First: "Miss",
		Last:  "Moneypenny",
	}
	fmt.Printf("p2\t%#v\n", p2)
	p2.Speak("I'm 008.")
	p2.NewName("Jenny")
	fmt.Printf("%#v\n\n", p2)
}
