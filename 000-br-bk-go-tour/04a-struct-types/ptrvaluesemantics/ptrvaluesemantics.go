package ptrvaluesemantics

import "fmt"

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

// Pointer receiver
func (p *Person) RunsPtr() bool {
	fmt.Println("Printing from in RunsPtr")
	p.Age++
	return true
}

// IMPORTANT
// Notice that both value semantics & pointer semantics run both methods
// with value receivers and pointer receivers
// methods sets are important for determining interface implementation

func PtrValSem() {
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
}
