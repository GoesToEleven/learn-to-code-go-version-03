package main

import (
	"fmt"
	"runtime"
)

func main() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	fmt.Printf("Before allocation: %v MiB\n", mem.Alloc/1024)
	// mem.Alloc gives the number of bytes allocated and not yet freed by the application.
	// The first /1024 converts bytes to kilobytes (1 MiB = 1024 bytes).
	// The second /1024 converts kilobytes to megabytes (1 MiB = 1024 MiB).

	mySlice := make([]int, 0)
	runtime.ReadMemStats(&mem)
	fmt.Printf("After allocation: %v MiB\n", mem.Alloc/1024)

	mySlice = nil

	runtime.GC() // Force garbage collection
	runtime.ReadMemStats(&mem)
	fmt.Printf("After GC: %v MiB\n", mem.Alloc/1024)

	// use mySlice variable
	fmt.Println(len(mySlice))
}

/*
The "i" in MiB or KiB stands for "binary."
These units—KiB (Kibibyte), MiB (Mebibyte), GiB (Gibibyte), etc.
are based on binary (base-2) calculations. In these units:

- 1 KiB is \(2^{10}\) or 1024 bytes
- 1 MiB is \(2^{20}\) or 1024 KiB or \(1024 \times 1024\) bytes
- 1 GiB is \(2^{30}\) or 1024 MiB or \(1024 \times 1024 \times 1024\) bytes
- and so on...

In contrast, units like KB (Kilobyte), MB (Megabyte), GB (Gigabyte), etc.,
are often used in the context of base-10 calculations, especially in fields
like storage and data transfer rates. In these units:

- 1 KB is \(10^3\) or 1000 bytes
- 1 MB is \(10^6\) or 1000 KB or \(1000 \times 1000\) bytes
- 1 GB is \(10^9\) or 1000 MB or \(1000 \times 1000 \times 1000\) bytes
- and so on...

However, it's important to note that the usage of KB, MB, GB, etc., can be ambiguous,
especially in the context of computing and data storage. Historically, many systems
and people have used these to refer to binary-based sizes, essentially using them
in place of KiB, MiB, GiB, etc. For example, many operating systems display file sizes
in "MB" or "GB," but they calculate those sizes using binary multiples (i.e., based on 1024 instead of 1000).

The introduction of the "i" notation aims to clear up this ambiguity. With this notation,
it's clear whether the number is calculated based on binary (base-2) or decimal (base-10) multiples.
*/
