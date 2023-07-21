package main

import "fmt"

type person struct {
	age int
}

func main() {
	people := make([]person, 2)

	p1 := &people[1]
	fmt.Printf("%p \n", p1)
	fmt.Printf("%p \n", &people[1])
	p1.age++
	fmt.Println(p1.age)
	fmt.Println(people[1].age)

	// Add a new person
	people = append(people, person{})

	// Add another year for p1
	p1.age++
	fmt.Printf("%p \n", p1)
	fmt.Printf("%p \n", &people[1])
	fmt.Println(people[1].age)
}
