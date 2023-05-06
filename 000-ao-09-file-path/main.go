package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var counter int
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.Contains(path, ".go") || strings.Contains(path, ".exe") {
			return nil
		}
		replace(path)
		counter++
		return nil
	})
	fmt.Printf("files changed: %v", counter)
}

// path is the filepath to the file
func replace(path string) {
	// open original file
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer f.Close()

	//read the file

	fmt.Println("path\t", path)
	fmt.Println("name\t", f.Name())
	// backslash must be escaped
	// https://go.dev/ref/spec#Rune_literals
	xs := strings.Split(f.Name(), "\\")
	fmt.Println(xs)
	fmt.Println("last element\t", xs[len(xs)-1])
	fmt.Print("\n--------------------\n")
}
