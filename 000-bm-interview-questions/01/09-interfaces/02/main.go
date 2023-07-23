package main

import "fmt"

type printer interface {
	print()
}

type cannon struct {
	name string
}

func (c cannon) print() {
	fmt.Printf("Printer Name: %s\n", c.name)
}

type epson struct {
	name string
}

func (e *epson) print() {
	fmt.Printf("Printer Name: %s\n", e.name)
}

func main() {
	c := cannon{"c-1"}
	e := epson{"e-2"}

	printers := []printer{
		c,
		&e,
	}
	fmt.Printf("e in printers scope - \t %p \n", printers[1])
	fmt.Printf("e in main scope - \t %p \n", &e)

	c.name = "c-2"
	e.name = "e-2"

	for _, v := range printers {
		doIt(c)
		doIt(&e)
		v.print()
		fmt.Println("^^^^^")
	}
}
func doIt(p printer) {
	p.print()
}
