package main

import "fmt"

const animalInfo string = `This animal's info is`

type dog struct {
	first string
	age   int
}

type cat struct {
	first string
	age   int
}

type animal interface {
	Info() string
}

func (c cat) Info() string {
	return fmt.Sprintf("%s %v", animalInfo, c)
}

func (d dog) Info() string {
	return fmt.Sprintf("%s %v", animalInfo, d)
}

func prtInfoInterface(aa []animal) {
	for i, a := range aa {
		fmt.Printf("%d - %s \n", i, a.Info())
	}
}

func main() {
	d1 := dog{"Rover", 7}
	d2 := dog{"Rufus", 8}
	c1 := cat{"Fluffy", 42}
	c2 := cat{"Buffy", 43}

	// using interface
	xa := []animal{d1, d2, c1, c2}
	prtInfoInterface(xa)

	// using generics
	xd := []dog{d1, d2}
	xc := []cat{c1, c2}
	prtInfo[dog](xd)
	prtInfo(xc) // type inference here
}

// generics

func info[T dog | cat](t T) string {
	return fmt.Sprintf("%s %v", animalInfo, t)
}

func prtInfo[T dog | cat](tt []T) {
	for i, t := range tt {
		fmt.Printf("%d - %s \n", i, info[T](t))
	}
}
