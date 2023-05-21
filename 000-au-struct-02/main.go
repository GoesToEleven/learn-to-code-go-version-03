package main

import "fmt"

type Person struct {
	First string
	Last  string
}

func main() {
	p := NewPerson("James", "Bond")
	fmt.Printf("%#v", p)
}

func NewPerson(f, l string) *Person {
	np := Person{f, l}
	return &np
}

// different types of New functions you can create
// API - pointer semantics
// data - value semantics
// func NewPerson(f, l string) Person  {}
// func NewPerson(f, l string) *Person {}
// func NewPerson(f, l string) (Person, error)  {}
// func NewPerson(f, l string) (*Person, error)  {}

//go build -gcflags=-m
// -- shows what escapes to the heap
//go clean
// -- removes your executable

/*
memory can be allocated on
-- THE STACK
---- function stack; self-cleaning;
---- memory allocations for variables used in a fuction scope
---- are released when the function completes
-- THE HEAP
---- garbage collector (gc) has to clean the heap
---- memory allocations on the heap outlive functions
-- ESCAPE ANALYSIS
---- compiler looks to see if a variable in a function will outlive the function
---- if so, memory allocation on THE HEAP
---- this happens when you return a pointer (memory address) from a function
*/
