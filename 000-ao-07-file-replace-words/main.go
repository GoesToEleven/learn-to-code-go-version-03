package main

import (
	"bytes"
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
	// read the file
	xb1, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	// replace bytes
	type replace string
	type with string
	m := map[replace]with{
		"wanna": "want to",
		"gonna": "going to",
	}
	xb2 := xb1
	for k, v := range m {
		xb2 = bytes.Replace(xb2, []byte(k), []byte(v), -1)
	}

	// new file name (nfn)
	// backslash must be escaped
	// https://go.dev/ref/spec#Rune_literals
	xs := strings.Split(path, "\\")
	nfn := xs[len(xs)-1]

	//write to new file
	if err = os.WriteFile(nfn, xb2, 0666); err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(nfn)
	fmt.Println("-----------------------")
}
