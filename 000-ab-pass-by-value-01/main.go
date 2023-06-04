package main

import "fmt"

//struct
type person struct {
	first string
	age   int
}

func main() {
	//slice
	xi := []int{1, 2, 3}
	fmt.Println(xi)
	changeSlice(xi, 1)
	fmt.Println(xi)

	//map
	m := map[string]int{
		"dog":  1,
		"cat":  2,
		"bird": 3,
	}
	fmt.Println(m)
	changeMap(m, "cat")
	fmt.Println(m)

	// value
	n := 42
	fmt.Println(n)
	changeVal(n)
	fmt.Println(n)
	// with pointer
	changeVal2(&n)
	fmt.Println("again ", n)

	//struct
	p1 := person{
		first: "Todd",
		age:   45,
	}

	fmt.Println(p1)
	changePerson(p1)
	fmt.Println(p1)
	changePerson2(&p1)
	fmt.Println(p1)
	p1.changePerson3()
	fmt.Println(p1)
	p1.changePerson4()

	//new person
	p2, err := newPerson("James", 42)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n", p2)
	}
}

func changeSlice(a []int, b int) {
	a[b] *= 2
}

func changeMap(c map[string]int, d string) {
	c[d] *= 2
}

func changeVal(e int) {
	e *= 2
}

func changeVal2(f *int) {
	*f *= 2
}

func changePerson(p person) {
	p.age = 50
}

func changePerson2(p *person) {
	p.age = 50
}

func (p *person) changePerson3() {
	p.age = 52
}

func (p person) changePerson4() {
	//won't work
	//p.age = 54
	fmt.Printf("%#v\n",p)
}

func newPerson(s string, n int) (*person, error) {
	if s == "" {
		return nil, fmt.Errorf("person cannot be empty")
	}
	if n <= 0 {
		return nil, fmt.Errorf("age must be greater than zero")
	}
	p := person{s, n}
	return &p, nil
}
/*
The method set of a type determines the methods that can be called 
on an operand of that type. 

defined type T
The method set of a defined type T consists of all methods declared with 
receiver type T.

pointer to a defined type T
The method set of a pointer to a defined type T 
(where T is neither a pointer nor an interface) 
is the set of all methods declared with 
receiver *T or T.

*/
// https://go.dev/ref/spec#Method_sets
