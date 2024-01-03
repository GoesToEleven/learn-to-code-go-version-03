package main

import "fmt"

func main() {
	// array literal
	dogs := [3]string{"fifi", "rover", "fred"}
	fmt.Println(dogs)
	fmt.Println(dogs[0])
	fmt.Println(dogs[1])
	fmt.Println(dogs[2])
	
	// array zero value
	var cats [4]string
	fmt.Println(cats)
	
	// array assignment by index position
	cats[0] = "fluffy"
	cats[1] = "puffy"
	cats[2] = "hissy"
	cats[3] = "Bucky"
	fmt.Println(cats)
	
	// array size is part of type
	fmt.Println("---------- SIZE")
	fmt.Printf("%T \n", dogs)
	fmt.Printf("%T \n", cats)
	// dogs = cats // doesn't work
	
	// for loop
	fmt.Println("---------- FOR LOOP")
	for i :=0; i < len(dogs); i++ {
		fmt.Println(dogs[i])
	}
	
	// for range clause
	fmt.Println("---------- FOR RANGE")
	for i, v := range dogs {
		fmt.Println(i, v)
	}
	
	// for range show just the value using the blank identifier (underscore character)
	fmt.Println("---------- FOR RANGE blank identifier")
	for _, v := range dogs {
		fmt.Println(v)
	}

	// for range show just the index
	fmt.Println("---------- FOR RANGE just index")
	for i := range dogs {
		fmt.Println(i, dogs[i])
	}

	// CONTIGUOUS MEMORY
	// address comparision: RANGE VALUE vs. INDEX POSITION VALUE
	fmt.Println("---------- FOR RANGE contiguous memory")
	for i, v := range dogs {
		fmt.Printf("%d \t %v \t %p \t %p \n", i, v, &v, &dogs[i])
	}

	// SEQUENTIAL CODE
	fmt.Println("-------- psychologists")

	// The for range cause seems to be like a function
	// It receives the data structure to range over
	// and then it ranges over it
	
	// range over a VALUE
	psychologists := [5]string{"Alpert", "C.Jung", "Piaget", "Rogers", "Skinner"}
	for i, v := range psychologists {
		psychologists[1] = "MASLOW"
		fmt.Printf("%d \t %s \t\t\t %s \n", i, v, psychologists[i])
	}
	fmt.Println(psychologists)
	
	// range over a POINTER
	psychologists = [5]string{"Alpert", "C.Jung", "Piaget", "Rogers", "Skinner"}
	for i, v := range &psychologists {
		psychologists[1] = "MASLOW"
		fmt.Printf("%d \t %s \t\t\t %s \n", i, v, psychologists[i])
	}
	fmt.Println(psychologists)

}
