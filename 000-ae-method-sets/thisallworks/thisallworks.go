package thisallworks

import "fmt"

/*
The method set
of a defined type T
methods declared with
receiver type T.

The method set
of a pointer to a defined type T
(where T is neither a pointer nor an interface)
methods declared with
receiver *T or T.
*/


type animal struct {
	Name string
}

func (a animal) np() {
	fmt.Println("non-pointer method", a.Name)
}

func (pa *animal) p() {
	fmt.Println("pointer method", pa.Name)
}

func RunMe() {
	x := &animal{"Bird"}
	x.np()
	x.p()

	y := animal{"Dolphin"}
	y.np()
	y.p()
}
