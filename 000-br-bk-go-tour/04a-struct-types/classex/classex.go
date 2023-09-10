package classex

import "fmt"

type Person struct {
	First string
	Age   int
}

type SecretAgent struct {
	Person
	First string
	Age int
	LTK bool
}

func RunMe() {
	fmt.Println("-------------- This is from RunMe ------------")
	somethingInternal()

	p1 := Person{
		First: "Jenny",
		Age:   27,
	}
	fmt.Println(p1)

	sa1 := SecretAgent{
		Person: Person{
			First: "Joe",
			Age:   39,
		},
		First: "James",
		Age:   32,
		LTK: true,
	}
	fmt.Println(sa1)
	fmt.Println(sa1.Person.First)
	fmt.Println(sa1.Person.Age)
	fmt.Println(sa1.First)
	fmt.Println(sa1.Age)
	fmt.Println(sa1.LTK)

	fmt.Println("-------------- This is THE END from RunMe ------------")
}

func somethingInternal() {
	fmt.Println("this is internal")
}
