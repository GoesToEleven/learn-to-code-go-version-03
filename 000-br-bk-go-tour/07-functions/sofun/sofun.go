package sofun

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func CountWords() {
	// Open the file
	file, err := os.Open("./sofun/gatsby.txt")
	if err != nil {
		log.Fatalf("error opening file: %s \n", err)
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)

	// Count the words
	count := 0
	for scanner.Scan() {
		count++
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		log.Fatalf("error scanning file: %s \n", err)
	}

	// Print the word count
	fmt.Printf("Total words: %d\n", count)
}