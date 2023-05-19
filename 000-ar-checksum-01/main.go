package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	s, err := sha1Sum("./simpleFile.txt")
	if err != nil {
		log.Fatalf("error getting sha1: %s", err)
	}
	fmt.Println(s)
}

func sha1Sum(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", nil
	}
	defer file.Close()

	w := sha1.New()
	if _, err := io.Copy(w, file); err != nil {
		return "", err
	}
	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil
}

// cat simpleFile.txt | sha1sum
