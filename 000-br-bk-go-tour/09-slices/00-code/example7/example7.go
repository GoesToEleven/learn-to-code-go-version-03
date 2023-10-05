// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show the components of a slice. It has a
// length, capacity and the underlying array.
package main

import "fmt"

func main() {

	// Create a slice with a length of 5 elements and a capacity of 8.
	fruits := make([]int, 5, 8)
	fruits[0] = 42
	fruits[1] = 43
	fruits[2] = 44
	fruits[3] = 45
	fruits[4] = 46

	inspectSlice(fruits)
}

// inspectSlice exposes the slice header for review.
func inspectSlice(slice []int) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i, v := range slice {
		fmt.Printf("[%d] %p %d\n",
			i,
			&slice[i],
			v)
	}
}
