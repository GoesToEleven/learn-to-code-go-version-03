package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// append slices
// https://go.dev/play/p/ERb8SBXaHhR
// https://go.dev/play/p/n3j7L21Y2UH

func main() {
	var xb []byte
	var counter int

	// open files in current directory, and all sub-directories
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.Contains(path, "output.txt") || !strings.Contains(path, ".txt") {
			return nil
		}
		xb2, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("error: %s", err)
		}
		xb2 = bytes.Trim(xb2,"\xef\xbb\xbf") // remove BOM
		xb2 = append(xb2, []byte("\n"+path)...)
		xb2 = append(xb2, []byte("\n\n\n\n\n\n")...)
		xb = append(xb, xb2...)
		counter++
		return nil
	})

	// write to new file
	if err := os.WriteFile("output.txt", xb, 0666); err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("files added together: %v", counter)
}

/*
The byte order mark (BOM) is a piece of information used to signify 
that a text file employs Unicode encoding, while also communicating 
the text stream's endianness. The BOM is not interpreted as 
a logical part of the text stream itself, but is rather 
an invisible indicator at its head.

In computing, endianness is the order or sequence of bytes of a word 
of digital data in computer memory. Endianness is primarily expressed 
as big-endian or little-endian. 

A big-endian system stores the most significant byte of a word 
at the smallest memory address and the least significant byte 
at the largest.
*/