package main

import "fmt"

type Person struct {
	first string
	age   int
}

func (p Person) Speak() {
	fmt.Println("My name is", p.first, "and my age is", p.age)
}

func main() {
	p1 := Person{
		first: "Jenny",
		age:   27,
	}

	p1.Speak()

}

/*
Create a user defined struct with
the identifier “person”
the fields:
first
age
attach a method to type person with
the identifier “speak”
the method should have the person say their name and age
create a value of type person
call the method from the value of type person
*/
