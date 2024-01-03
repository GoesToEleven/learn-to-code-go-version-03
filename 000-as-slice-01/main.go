package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("%#v\n\n", s)

	// len & cap are nil safe
	fmt.Printf("len: %#v\n", len(s))
	fmt.Printf("cap: %#v\n\n", cap(s))

	// you can compare a slice only to nil
	if s == nil {
		fmt.Printf("%v is a nil slice\n", s)
		fmt.Printf("%#v is a nil slice\n\n", s)
	}

	s2 := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("s2 = %#v\n\n", s2)

	// slicing a slice
	// [inclusive:exclusive]
	fmt.Println("************")
	fmt.Printf("s2[1:4] = %#v\n", s2[1:4])
	fmt.Printf("s2[1:] = %#v\n", s2[1:])
	fmt.Printf("s2[:4] = %#v\n", s2[:4])
	fmt.Printf("s2[0:4] = %#v\n", s2[0:4])
	fmt.Printf("s2[:] = %#v\n\n", s2[:])

	// doesn't work - panic
	// fmt.Printf("s2[:100] = %#v\n", s2[:100])

	// append
	fmt.Printf("len(s2): %#v\n", len(s2))
	fmt.Printf("cap(s2): %#v\n", cap(s2))
	s2 = append(s2, 42)
	fmt.Printf("s2 with 42 appended: %#v\n", s2)
	fmt.Printf("len(s2) with 42 appended: %#v\n", len(s2))
	fmt.Printf("cap(s2) with 42 appended: %#v\n", cap(s2))
	fmt.Println(`
	Why did cap increase?
	A new underlying array was created
	Values from the previous underlying array were copied into it
	the previous underlying array was destroyed
	takeaway:
	if you know in advance how many items you are going to store in a slice
	use make to make all of your allocations at once
	`)

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

	s3 := s2[1:4]
	fmt.Printf("s3: %#v\n", s3)
	s3 = append(s3, 43)
	fmt.Printf("s3 with 43 appended to s3: %#v\n", s3)
	fmt.Printf("s2 with 43 appended to s3: %#v\n\n", s2)

	// make
	// https://pkg.go.dev/builtin#make
	s4 := make([]int, 10)
	fmt.Printf("len: \t %#v\n", len(s4))
	fmt.Printf("cap: \t %#v\n", cap(s4))
	fmt.Printf("s4: \t %#v\n", s4)
	s4[4] = 4
	fmt.Printf("s4: \t %#v\n\n", s4)

	// add values to s4
	for i := 0; i < 10; i++ {
		s4[i] = i
	}

	s5 := s4[2:5]
	fmt.Printf("s4: \t %#v\n", s4)
	fmt.Printf("s5: \t %#v\n\n", s5)
	//
	fmt.Printf("len s4: %#v\n", len(s4))
	fmt.Printf("cap s4: %#v\n", cap(s4))
	fmt.Printf("len s5: %#v\n", len(s5))
	fmt.Printf("cap s5: %#v\n", cap(s5))
	fmt.Println(`
	Why is the cap 8 for s5?
	Because the slice data structure has a pointer that points to an underlying array
	and the pointer for s5 is pointing to the position in the array of the slice's first value 
	which is the int 2 in the third index position of the underlying array
	and from that third index position to the end of the underlying array is 8 index positions
	`)

	// copy
	// https://pkg.go.dev/builtin#copy
	s6 := make([]int, len(s5))
	copy(s6, s5)
	fmt.Printf("s5: \t %#v\n", s5)
	fmt.Printf("s6: \t %#v\n\n", s6)
	fmt.Printf("len s5: %#v\n", len(s5))
	fmt.Printf("cap s5: %#v\n", cap(s5))
	fmt.Printf("len s6: %#v\n", len(s6))
	fmt.Printf("cap s6: %#v\n", cap(s6))
	s5[1] = 42
	fmt.Printf("s5: \t %#v\n", s5)
	fmt.Printf("s6: \t %#v\n\n", s6)

	fmt.Println(`
as mentioned before
if you know in advance how many items you are going to store in a slice
use make to make all of your allocations at once

Slice: The size specifies the length. The capacity of the slice is
equal to its length. A second integer argument may be provided to
specify a different capacity; it must be no smaller than the
length. For example, make([]int, 0, 10) allocates an underlying array
of size 10 and returns a slice of length 0 and capacity 10 that is
backed by this underlying array.

https://pkg.go.dev/builtin#make
	`)

	// var s7 []int
	s7 := make([]int, 0, 1_000)
	fmt.Printf("%s\t%s\n", "len", "cap")
	for i := 0; i < 10; i++ {
		s7 = appendInt(s7, i)
	}
	fmt.Printf("len: %d \ncap: %d \n\n", len(s7), cap(s7))
}

func appendInt(s []int, v int) []int {
	i := len(s)
	fmt.Printf("%d\t%d\n", len(s), cap(s))
	if len(s) < cap(s) {
		s = s[:len(s)+1]
	} else {
		fmt.Printf("reallocate %d-->%d\n", len(s), 2*len(s)+1)
		s2 := make([]int, 2*len(s)+1)
		copy(s2, s)
		s = s2[:len(s)+1]
	}
	s[i] = v
	return s
}
