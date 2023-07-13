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
	d1 := dog{"Rover", 7}
	d2 := dog{"Rufus", 8}
	c1 := cat{"Fluffy", 42}
	c2 := cat{"Buffy", 43}

	// we can use a function that is more GENERIC using generics
	xd := []dog{d1, d2}
	prtInfo[dog](xd)

	xc := []cat{c1, c2}
	prtInfo[cat](xc)
}

// we can use a function that is more GENERIC using generics

func info[T dog | cat](t T) string {
	return fmt.Sprintf("%s %v", animalInfo, t)
}

func prtInfo[T dog | cat](tt []T) {
	for i, t := range tt {
		fmt.Printf("%d - %s \n", i, info[T](t))
	}
}
