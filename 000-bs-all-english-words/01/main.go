package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	// Open the folder and read its contents
	files, err := ioutil.ReadDir("files")
	if err != nil {
		log.Fatalln("Error reading directory:", err)
		return
	}

	// Read the words in each file
	m, err := readWords(files)
	if err != nil {
		log.Fatalln("Error reading files:", err)
		return
	}

	// Print the quantity of words from all of the files
	fmt.Println(len(m))

	// Sort the words
	xw := sortWords(m)

	// Write the words to a txt file
	writeWords(xw)

}

// readWords reads all of the words
func readWords(ff []fs.FileInfo) (map[string]int, error) {
	// Create a map for this specific file
	words := make(map[string]int)

	// Loop through each file in the folder
	for _, file := range ff {
		filename := file.Name()

		// Process only .txt files
		if strings.HasSuffix(filename, ".txt") {

			// Open the file for reading
			f, err := os.Open("files/" + filename)
			if err != nil {
				fmt.Println("Error opening file:", err)
				return words, err
			}

			// Create a new scanner and scan the file line by line
			scanner := bufio.NewScanner(f)

			// Split all of the text into words
			scanner.Split(bufio.ScanWords)

			// Loop through all of the words
			for scanner.Scan() {
				word := scanner.Text()
				xw := strings.Split(word, ",")
				for _, w := range xw {
					w = cleanWord(w)
					// Add word to map
					if w != "" {
						words[w]++
					}

				}
			}

			// Check for errors during scanning
			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading file:", err)
				return words, err
			}

			// Close the file
			f.Close()
		}
	}
	return words, nil
}

// SortWords puts the words into a slice and sorts them
func sortWords(m map[string]int) []string {
	// Extract words into a slice
	words := make([]string, 0, len(m))
	for k := range m {
		words = append(words, k)
	}

	// Sort the words
	sort.Strings(words)

	return words
}

func writeWords(xw []string) {
	// Create or open a file
	file, err := os.Create("all-words.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a buffered writer for better performance
	writer := bufio.NewWriter(file)

	// Write each word to the file
	for _, word := range xw {
		_, err := writer.WriteString(word + "\n")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	// Flush the buffer to ensure all data gets written to file
	writer.Flush()
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
