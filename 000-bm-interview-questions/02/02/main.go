package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Printf("%p \t %T \n", &x, &x)
	fmt.Println(*&x)

	sports := make([]string, 5)
	sports[0] = "ski"
	sports[1] = "surf"
	sports[2] = "swim"
	sports[3] = "sail"
	sports[4] = "sumo wrestling"

	for i, v := range sports {
		fmt.Printf("%v - %p \t %v - %p \n", v, &v, sports[i], &sports[i])
	}
}
