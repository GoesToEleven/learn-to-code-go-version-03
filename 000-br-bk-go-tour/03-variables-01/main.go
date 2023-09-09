package main

import (
	"fmt"
	"unsafe"
)

// Note that I've used the `unsafe` package here to determine the size,
// unsafe is not recommended for everyday Go programming
// but it is okay for this kind of introspection

func main() {
	fmt.Printf("\n---------------- BOOL ----------------\n")
	var b bool
	fmt.Println("SIZE OF BOOL:", unsafe.Sizeof(b), "byte")

	fmt.Printf("\n---------------- INT -----------------\n")
	var c int
	fmt.Printf("SIZE OF INT (on this machine) \t %d bytes \n", unsafe.Sizeof(c))
	var d1 int8
	fmt.Printf("SIZE OF INT8 \t \t \t %d bytes \n", unsafe.Sizeof(d1))
	var d2 int16
	fmt.Printf("SIZE OF INT16 \t \t \t %d bytes \n", unsafe.Sizeof(d2))
	var d3 int32
	fmt.Printf("SIZE OF INT32 \t \t \t %d bytes \n", unsafe.Sizeof(d3))
	var d4 int64
	fmt.Printf("SIZE OF INT64 \t \t \t %d bytes \n", unsafe.Sizeof(d4))

	fmt.Printf("\n---------------- WORD ----------------\n")
	fmt.Println(`In computing, the term "word" refers to a fixed-sized piece of data 
that a machine can process as a single unit. The size of a word varies 
depending on the architecture of the machine. For example:

- On a 32-bit system, a word is usually 32 bits, or 4 bytes.
- On a 64-bit system, a word is usually 64 bits, or 8 bytes.

The word size determines the size of various data types that are native to the processor, 
including pointers, integers, and floating-point numbers. 

Understanding word size can be important for performance tuning, 
as accessing data in chunks that align with the machine's word size 
is generally more efficient than accessing data in smaller or larger chunks.`)

	fmt.Printf("\n---------------- STRING --------------\n")
	var f string
	fmt.Println("SIZE OF STRING (zero value: pointer is nil):", unsafe.Sizeof(f), "byte")
	fmt.Printf("str data structure \t 0x%x\n", (*[2]uintptr)(unsafe.Pointer(&f)))
	fmt.Println("")
	f = "H🚀"
	fmt.Println("SIZE OF STRING (with pointer to underlying bytes):", unsafe.Sizeof(f), "byte")
	fmt.Printf("str data structure \t 0x%x\n", (*[2]uintptr)(unsafe.Pointer(&f)))
	fmt.Printf("len of string \t\t %d \n", len(f))

	/*
		type stringStruct struct {
		str unsafe.Pointer
		len int
		}

		// source: src/runtime/string.go
	*/

	/*
		In Go, the UNDERLYING BYTES of a string are not formally stored in a Go array (`[N]byte`),
		but they are STORED IN A CONTIGUOUS BLOCK OF MEMORY, much like how an array works at the memory level.
		The `str unsafe.Pointer` field in the `stringStruct` points to the first byte of this contiguous block of memory.

		Since strings in Go are UTF-8 encoded,
		each character in the string may take up 1 to 4 bytes.
		This contiguous block of memory will be exactly `len` bytes long,
		as indicated by the `len` field in the `stringStruct`.

		For example, if you have a Go string `s := "hello"`,
		a contiguous block of 5 bytes will be allocated in memory to hold the bytes
		that represent the characters 'h', 'e', 'l', 'l', 'o'. The `str` pointer will point to the first of these bytes,
		and `len` will be 5.

		Because strings in Go are immutable,
		this memory block will not be modified once it is created.
		Any string manipulation operations that you perform will result
		in a new block of memory being allocated for the new string.

		It's worth noting that due to the immutability of strings and Go's optimizations,
		substrings often share the same underlying block of memory with their parent strings to minimize memory allocations.
		For instance, if you slice a string `s` to get a substring `t`, both `s` and `t` might point to different offsets
		within the same block of memory, provided no new modifications are needed to create `t`.
	*/

	fmt.Printf("\n---------------- STRING (runes) ------\n")

	// see each rune, aka, character
	for _, runic := range f {
		fmt.Printf("%d \t %x \t %s \n", runic, runic, string(runic))
	}

	fmt.Printf("\n---------------- STRING (bytes) ------\n")
	// see bytes in each rune
	for _, runic := range f {
		xb := []byte(string(runic))
		for _, b := range xb {
			fmt.Printf("%d \t \t %s \n", b, string(runic))
		}
	}

	fmt.Printf("\n---------------- SLICE ---------------\n")
	var g []int
	fmt.Println("SIZE OF SLICE (zero value)", unsafe.Sizeof(g), "byte")
	fmt.Printf("slice data structure \t 0x%x\n", (*[3]uintptr)(unsafe.Pointer(&g)))
	fmt.Println("")
	g = []int{1, 4, 9, 42}
	fmt.Println("SIZE OF SLICE (with pointer to underlying array)", unsafe.Sizeof(g), "byte")
	fmt.Printf("slice data structure \t 0x%x\n", (*[3]uintptr)(unsafe.Pointer(&g)))
	/*
		The Go programming language internally represents slices using a data structure with three fields:
		a pointer to the underlying array, a length, and a capacity.

		```go
		type slice struct {
			array unsafe.Pointer
			len   int
			cap   int
		}
		```

		// source: src/runtime/slice.go

		- `array`: A pointer to the first element of the array that holds the data for the slice.
		It points to an array in memory that is at least as large as the specified capacity (`cap`).

		- `len`: An integer representing the current length of the slice,
		i.e., the number of elements it contains. The `len` should be less than or equal to `cap`.

		- `cap`: An integer representing the capacity of the slice,
		i.e., the maximum number of elements that can be stored in the underlying array without reallocating memory.

		IMPORTANT
		When you slice an existing slice (or array),
		Go creates a new slice data structure with a new `len` and potentially a new `cap`,
		but both slices share the same underlying array unless the slice operation requires a reallocation of the array.
		This makes slice operations very efficient, but programmers must be cautious when multiple slices share the same
		underlying array, as changes to one can affect the others.
	*/

}
