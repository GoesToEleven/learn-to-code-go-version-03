package main

import "fmt"

type printer interface {
	print()
}

type cannon struct {
	name string
}

type epson struct {
	name string
}

type hp struct {
	name string
}

func (c cannon) print() {
	fmt.Printf("Printer Name: %s\n", c.name)
}

func (e *epson) print() {
	fmt.Printf("Printer Name: %s\n", e.name)
}

func (h *hp) print() {
	fmt.Printf("Printer Name: %s\n", h.name)
}

func main() {
	c := cannon{"c-1"}
	e := &epson{"e-1"}
	h := hp{"h-1"}

	fmt.Println(c.name)
	fmt.Println(e.name)
	fmt.Println(h.name)
	c.print()
	e.print()
	h.print()
	doIt(c)
	doIt(e)
	// doIt(h)
	// doesn't implement interface
	// value of type T can only have
	// methods with receivers of type T to implement an interface
	fmt.Println("************")

	printers := []printer{
		c,
		e,
		&h,
	}
	doIt(printers[0])
	doIt(printers[1])
	doIt(printers[2]) // does implement interface now
	fmt.Println("in printers h is stored at", &printers[2])
	fmt.Printf("in main h is stored at %p \n", &h)
	fmt.Println("************")

	c.name = "c-2"
	e.name = "e-2"
	h.name = "h-2"
	fmt.Println(c.name)
	fmt.Println(e.name)
	fmt.Println(h.name)
	c.print()
	e.print()
	h.print()
	doIt(c)
	doIt(e)
	// doIt(h)
	fmt.Println("************")

	for _, v := range printers {
		v.print()
		doIt(v)
	}
}

func doIt(p printer) {
	p.print()
}
