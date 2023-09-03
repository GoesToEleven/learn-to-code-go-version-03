package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	// Read all valid words
	allWords, err := ReadWords("all-words.txt")
	if err != nil {
		fmt.Println("Error reading all-words.txt:", err)
		return
	}

	// Open the file to check
	fileToCheck, err := os.Open("check-spelling.txt")
	if err != nil {
		fmt.Println("Error opening check-spelling.txt:", err)
		return
	}
	defer fileToCheck.Close()

	checkSpelling(fileToCheck, allWords)
}

// ReadWords reads a list of words from a file and returns a set.
func ReadWords(filename string) (map[string]bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	words := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		words[word] = true
	}
	return words, scanner.Err()
}

func checkSpelling(f *os.File, allWords map[string]bool) {

	// Create a scanner and read the file line by line.
	scanner := bufio.NewScanner(f)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())
		wordsInLine := strings.Fields(line)

		for _, word := range wordsInLine {
			word = cleanWord(word)

			_, exists := allWords[word]
			if word != "" && !exists {
				fmt.Printf("Line %d: Spelling error in word %s\n", lineNumber, word)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading check-spelling.txt:", err)
	}
}

// CleanWord removes digits from the word
func cleanWord(word string) string {

	word = strings.ToLower(word)
	// Remove leading and trailing undesired characters
	word = strings.Trim(word, "*.,!?\"'()[]{}:;#&+-/=$%<>@_|~")
	// Remove an leading and trailing spaces
	word = strings.TrimSpace(word)

	// Discard word with undesired characters
	xs := []string{"\a", "\b", "\f", "\n", "\r", "\t", "\v", "\\", "\"", "-", ".", "=", "'", "`", "?", "$", "’", "“", "‘", "–"}
	for _, v := range xs {
		if strings.Contains(word, v) {
			return ""
		}
	}

	// Discard word with numeric digits
	for _, r := range word {
		if unicode.IsDigit(r) {
			return ""
		}
	}

	return word
}
