package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {

	v1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Ford",
		model: "Mustang",
		doors: 4,
		color: "Blue",
	}

	v2 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "Toyota",
		model: "Tundra",
		doors: 2,
		color: "White",
	}

	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Println(v1.electric, v1.make, v1.model)
	fmt.Println(v2.electric, v2.make, v2.model)

	fmt.Println(v1.engine.electric, v1.make, v1.model)
	fmt.Println(v2.engine.electric, v2.make, v2.model)
}

/*
Create a type engine struct, and include this field
electric bool
Create a type vehicle struct, and include these fields
engine
make
model
doors
color
Create two VALUES of TYPE vehicle
use a composite literal
Print out each of these values.
Print out a single field from each of these values.

*/
