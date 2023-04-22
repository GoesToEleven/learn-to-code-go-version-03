package main

import "fmt"

type person struct {
	first string
}

func (p person) String() string {
	return fmt.Sprint("My name is ", p.first)
}

func main() {
	p1 := person{"James"}
	fmt.Println(p1)
}

/*
type Stringer

type Stringer interface {
	String() string
}

Stringer is implemented by any value that has a String method, 
which defines the “native” format for that value. The String method 
is used to print values passed as an operand to any format 
that accepts a string or to an unformatted printer such as Print.

https://pkg.go.dev/fmt#Stringer
*/