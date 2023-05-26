package main

import "fmt"

func main() {
	xs := []string{"hello", "world"}
	fmt.Println(xs)
}

/*
	slice is a data structure with three elements:
	-- len
	-- cap
	-- pointer to an underlying array

	type slice struct {
		array unsafe.Pointer
		len   int
		cap   int
	}
*/
// src/runtime/slice.go
// google: "golang pkg runtime" then click on the "slice.go" file at the bottom of the page

/*
"Arrays have their place, but they’re a bit inflexible,
so you don’t see them too often in Go code.
Slices, though, are everywhere.
They build on arrays to provide
great power and convenience."
~ Go Slices: usage and internals - Go Blog - Andrew Gerrand
*/
// https://go.dev/blog/slices-intro

/*
	slice is a data structure with three elements:
	-- len
	-- cap
	-- pointer to an underlying array

	type slice struct {
		array unsafe.Pointer
		len   int
		cap   int
	}
*/
// src/runtime/slice.go
// google: "golang pkg runtime" then click on the "slice.go" file at the bottom of the page
