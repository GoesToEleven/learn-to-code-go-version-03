package main

import "fmt"

type animal interface {
	speaks()
}

type dog struct {
	saying   string
	isMammal bool // duplication is cheaper than the wrong abstraction
}

func (d *dog) speaks() {
	fmt.Println(d.saying)
}

type cat struct {
	saying   string
	isMammal bool // duplication is cheaper than the wrong abstraction
}

func (c *cat) speaks() {
	fmt.Println(c.saying)
}

func main() {
	d1 := &dog{"woof", true}
	c1 := &cat{"meow", true}
	d2 := &dog{"woof woof", true}
	c2 := &cat{"meeeoooooowwwwwww", true}

	xa := []animal{d1, c1, d2, c2}

	for _, a := range xa {
		a.speaks()
	}
}
