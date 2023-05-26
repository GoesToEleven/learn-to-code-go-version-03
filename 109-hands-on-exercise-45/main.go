package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	
	x = append(x, 52)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

/*

Follow these steps:
start with this slice
x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
append to that slice this value
52
print out the slice
in ONE STATEMENT append to that slice these values
53
54
55
print out the slice
append to the slice this slice
y := []int{56, 57, 58, 59, 60}
print out the slice

*/
