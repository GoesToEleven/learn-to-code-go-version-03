package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	data := []byte("Hello, world!")
	r := bytes.NewReader(data)

	// You can now read from 'r' as if it were any other io.Reader
	io.Copy(os.Stdout, r)
}

/*
The `bytes.NewReader` function in Go creates a new instance of `bytes.Reader` that reads from the provided byte slice. It is often used when you want to treat a byte slice as a read-only stream of bytes.

In this code, `bytes.NewReader(data)` creates a new `bytes.Reader` from the `data` byte slice. `io.Copy(os.Stdout, r)` then reads from this `Reader` and writes to the standard output, printing "Hello, world!" to the console.

A `bytes.Reader` satisfies the `io.Reader`, `io.ReaderAt`, `io.WriterTo`, `io.Seeker`, `io.ByteScanner`, and `io.RuneScanner` interfaces from the Go standard library, so it can be used wherever these interfaces are expected.
*/
