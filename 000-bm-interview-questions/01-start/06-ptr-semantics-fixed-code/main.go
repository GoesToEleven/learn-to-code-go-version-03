package main

import "fmt"

type person struct {
	age int
}

func main() {
	people := make([]person, 2)

	people[1].age++
	fmt.Println(people[1].age)

	// Add a new person
	people = append(people, person{})

	// Add another year for p1
	people[1].age++
	fmt.Println(people[1].age)
}
