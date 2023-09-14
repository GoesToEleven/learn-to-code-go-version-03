package main

import "fmt"

func main() {
	fmt.Printf("VALUE SEMANTICS \n \n")

	p1 := Person{"James", 32}
	fmt.Println(p1.Describe())
	fmt.Println(p1.IsAdult())

	fmt.Println()

	fmt.Println(p1.Age)
	fmt.Println(p1.RunsPtr())
	fmt.Println(p1.Age)

	fmt.Printf("\nPOINTER SEMANTICS \n \n")

	p2 := &Person{"Jenny", 27}
	fmt.Println(p2.Describe())
	fmt.Println(p2.IsAdult())

	fmt.Println()

	fmt.Println(p2.Age)
	fmt.Println(p2.RunsPtr())
	fmt.Println(p2.Age)

	p3 := &Person{Age: 55}
    p4 := p3        // Both point to the same data
	fmt.Println(p3, p4)
    p4.Age = 57     // Affects both p1 and p2
	fmt.Println(p3, p4)
}

type Person struct {
	Name string
	Age  int
}

// Value receiver
func (p Person) Describe() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

// Value receiver
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

func (p *Person) RunsPtr() bool {
	fmt.Println("Printing from in RunsPtr")
	p.Age++
	return true
}
