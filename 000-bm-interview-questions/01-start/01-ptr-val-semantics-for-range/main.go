package main

import "fmt"

func main() {
	sports := make([]string, 5)
	sports[0] = "ski"
	sports[1] = "surf"
	sports[2] = "swim"
	sports[3] = "sail"
	sports[4] = "sumo wrestling"
	
	fmt.Println("---------------")
	for _, v := range sports {
		fmt.Println(v)
	}

	fmt.Println("---------------")
	for i := range sports {
		fmt.Println(sports[i])
	}

	// ILLUSTRATED
	// comment each of these out - only run one at a time

	fmt.Println("---------------")
	for i, v := range sports {
		sports[i] = "biking"
		fmt.Println(v)
	}
	fmt.Println(sports)

	/*
		
		fmt.Println("---------------")
		for i, v := range sports {
			sports[i] = "biking"
			fmt.Println(v)
		}

		
		fmt.Println("---------------")
		for i, v := range sports {
			sports[1] = "biking"
			if i == 1 {
				fmt.Println("sports 1: ", v)
			}
		}

		
		fmt.Println("---------------")
		for i := range sports {
			sports[1] = "biking"
			if i == 1 {
				fmt.Println("sports 1: ", sports[1])
			}
		}
	*/
}
