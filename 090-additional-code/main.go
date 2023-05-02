package main

import "fmt"

func main() {
	x := "James Bond"
	if x == "James Bond" {
		fmt.Println(x)
	}

	y := "Moneypenny"
	if y == "Moneypenny" {
		fmt.Println(y)
	} else if y == "James Bond" {
		fmt.Println("BONDDONBONDONBOND", y)
	} else {
		fmt.Println("neither")
	}

	switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("should print")
	}

	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to hawaii!")
	case "wingsuit flying":
		fmt.Println("Living dangerously!")
	}
}
