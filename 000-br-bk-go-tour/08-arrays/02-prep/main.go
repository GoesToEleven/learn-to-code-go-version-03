package main

import "fmt"

func main() {
	var animals [5]string
	animals[0] = "Aardvark"
	animals[1] = "Platypus"
	animals[2] = "Narwhal"
	animals[3] = "Tarsier"
	animals[4] = "Kakapo"

	fmt.Println("--------- RANGE LOOP ---------")
	// COPIES THE VALUE from the `animals` array to the variable `v`
	// `v` gets its own copy of the value
	for i, v := range animals {
		fmt.Println(i, v)
	}

	fmt.Println("--------- FOR LOOP ---------")
	// access values BY INDEX POSITION
	for i := 0; i < len(animals); i++ {
		fmt.Println(i, animals[i])
	}

	// changing v doesn't change `animals`
	fmt.Println("--------- v changed ---------")
	for _, v := range animals {
		v = "v changed"
		fmt.Println(v)
	}
	fmt.Println("ANIMALS", animals)

	// changing the value at an index position does change `animals`
	fmt.Println("--------- animals changed ---------")
	for i := range animals {
		animals[i] = "animals changed"
	}
	fmt.Println("ANIMALS", animals)
}

/*
A string is a two word data structure:

```go
type stringStruct struct {
    str unsafe.Pointer
    len int
}
```
src/runtime/string.go

The UNDERLYING BYTES of a string are not formally stored in a Go array (`[N]byte`),
but they are STORED IN A CONTIGUOUS BLOCK OF MEMORY, much like how an array works at the memory level.
The `str unsafe.Pointer` field in the `stringStruct` points to the first byte of this contiguous block of memory.

Go strings are immutable, meaning you can't change their underlying bytes once they've been created.
This is in contrast to slices and arrays, which are mutable.

A Go string can be thought of as a read-only slice of bytes with a fixed length.
The underlying bytes are stored in a way that's optimized for efficiency and performance by the Go runtime.
When you create a new string by slicing an existing one, for example, both strings will often share
the same underlying bytes in memory to reduce allocations and copying. This is safe because the bytes are immutable.

Although you can't change a string's underlying bytes,
you can obtain a slice of bytes representing the string's data using a conversion:

```go
s := "hello, world"
b := []byte(s)
```

After this conversion, b will be a slice that contains the same bytes as s, but unlike the string, this slice is mutable.

*/
