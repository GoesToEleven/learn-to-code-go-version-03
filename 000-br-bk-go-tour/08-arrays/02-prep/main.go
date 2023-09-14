package main

import "fmt"

func main() {
	// array literal
	dogs := [3]string{"labradoodle", "chihuahua", "appenzeller"}
	fmt.Println(dogs)
	fmt.Println(dogs[0])
	fmt.Println(dogs[1])
	fmt.Println(dogs[2])

	// array zero value
	var cats [4]string
	fmt.Println(cats)
	fmt.Printf("%T \n", cats)

	// array assignment by index position
	cats[0] = "calico"
	cats[1] = "fuzzy"
	cats[2] = "round"
	cats[3] = "hissy"
	fmt.Println(cats)

	// array size is part of type
	// dogs = cats
	fmt.Println("AGAIN", dogs)

	// for loop
	for i := 0; i < len(dogs); i++ {
		fmt.Println(dogs[i])
	}

	// for range clause
	for i, v := range dogs {
		fmt.Println(i, v)
	}

	// for range show just the value using the blank identifier (underscore character)
	for _, v := range dogs {
		fmt.Println(v)
	}

	// for range show just the index
	for i := range dogs {
		fmt.Println(i, i, i, dogs[i])
	}

	// address comparision: RANGE VALUE vs. INDEX POSITION VALUE

	// SEQUENTIAL CODE
	fmt.Println("-------- psychologists")

	// The for range cause seems to be like a function
	// It receives the data structure to range over
	// and then it ranges over it

	// range over a VALUE
	psychologists := [5]string{"Alpert", "C.Jung", "Piaget", "Rogers", "Skinner"}
	for i, v := range psychologists {
		psychologists[1] = "MASLOW"
		fmt.Printf("%d \t %s \t \t \t %s \n", i, v, psychologists[1])
	}
	fmt.Println(psychologists)

	// range over a POINTER
	psychologists = [5]string{"Alpert", "C.Jung", "Piaget", "Rogers", "Skinner"}
	for i, v := range &psychologists {
		psychologists[1] = "MASLOW"
		fmt.Printf("%d \t %s \t \t \t %s \n", i, v, psychologists[1])
	}
	fmt.Println(psychologists)

}
