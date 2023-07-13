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

func main() {
	d := dog{"Rover", 7}
	c := cat{"Fluffy", 42}

	fmt.Println(c.cInfo())
	fmt.Println(d.dInfo())

	// we can use a function that is more GENERIC using generics

	fmt.Println(info[cat](c))
	fmt.Println(info(d)) // type inference here
}

func (c cat) cInfo() string {
	return fmt.Sprintf("%s %v", animalInfo, c)
}

func (d dog) dInfo() string {
	return fmt.Sprintf("%s %v", animalInfo, d)
}

// we can use a function that is more GENERIC using generics

func info[T dog | cat](t T) string {
	return fmt.Sprintf("%s %v", animalInfo, t)
}
