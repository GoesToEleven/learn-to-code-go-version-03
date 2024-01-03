// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how the for range has both value and pointer semantics.
package main

import "fmt"

func main() {

	// Using the value semantic form of the for range.
	friends := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	for _, v := range friends {
		friends = friends[:2]
		fmt.Printf("v[%s]\n", v)
	}
	fmt.Println(friends)

	fmt.Println()

	// Using the pointer semantic form of the for range.
	friends = []string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	for i, v := range friends {
		friends = friends[:2]
		if len(friends) > 2 {
			fmt.Printf("p[%s]\n", friends[i])
		}
		fmt.Println(i, v)
	}
}
