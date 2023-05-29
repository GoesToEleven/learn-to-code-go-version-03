package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	// Open a file
	f, err := os.Open("great-gatsby.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer f.Close()

	// The frequency of words in the file
	words, xs, err := freq(f)
	if err != nil {
		log.Fatalf("error from freq in main: %s", err)
	}

	// Display the word frequencies
	// for word, frq := range words {
	// 	fmt.Printf("%s \t\t %d\n", word, frq)
	// }

	// Sort the word frequencies
	pairs := sortWordFrequency(words)

	// Print the sorted results
	for _, pair := range pairs {
		fmt.Printf("%s \t\t %d\n", pair.Key, pair.Value)
	}

	// word with greatest frequency, and it's frequency
	w, n, err := maxWord(words)
	if err != nil {
		log.Fatalf("error with maxWord: %s\n", err)
	}
	fmt.Printf("%#v has a frequency %d", w, n)
	fmt.Printf("%#v", xs)
}

func freq(r io.Reader) (map[string]int, []string, error) {
	var i int
	xs := make([]string, 0, 1_000)
	fmt.Println("LENGTH of the slice",len(xs))
	fmt.Println("CAPACITY of the slice", cap(xs))
	// Create a map to store word frequencies
	wordFreq := make(map[string]int)

	// Create a scanner to read the file
	s := bufio.NewScanner(r)
	s.Split((bufio.ScanWords))

	// Read each word and update the word frequencies
	for s.Scan() {
		word := strings.ToLower(s.Text())
		wordFreq[word]++
		if i < 1_000 {
			word = strings.Replace(word, `"`, "", -1)
			word = strings.Replace(word, `“`, "", -1)
			word = strings.Replace(word, `”`, "", -1)
			xs = append(xs, word)
		}
		i++
	}
	if err := s.Err(); err != nil {
		return nil, nil, err
	}

	fmt.Println(xs)
	fmt.Println("LENGTH of the slice",len(xs))
	fmt.Println("CAPACITY of the slice", cap(xs))
	return wordFreq, xs, nil
}

// For sorting frequency of words
type Pair struct {
	Key   string
	Value int
}

// implement the Len, Less, and Swap methods on PairList
// to satisfy the sort.Interface interface.

type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	return p[i].Value > p[j].Value // Sort in descending order
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func sortWordFrequency(m map[string]int) PairList {
	// Convert the map to a pair list
	pairs := make(PairList, len(m))
	i := 0
	for key, value := range m {
		pairs[i] = Pair{key, value}
		i++
	}

	// Sort the pair list
	sort.Sort(pairs)

	return pairs
}

// word frequency in go
// https://chat.openai.com/share/6cb2b004-4cfa-4aad-a2d8-47b0eacd36dd

// sort maps in go
// turn them into structs
// get a slice of those structs
// implement the sort.Interface interface
// https://chat.openai.com/share/03a44e91-fc0d-4cdb-884a-c8acd8f440d8

func maxWord(m map[string]int) (string, int, error) {
	if len(m) == 0 {
		return "", 0, fmt.Errorf("empty map")
	}

	maxW := "" // word with max frequency
	maxF := 0  // max frequency of that word

	for k, v := range m {
		if v > maxF {
			maxW = k
			maxF = v
		}
	}
	return maxW, maxF, nil
}

/*
In the context of the `bufio` package in Go, a "token" refers to a sequence of characters that represents a meaningful unit of data within an input stream. The `bufio` package provides functionality for reading and writing data in buffered streams, and it includes a `Scanner` type that is designed to scan and tokenize input.

The `Scanner` type in the `bufio` package has a method called `Scan()`, which reads the next token from the input stream and advances the scanner to the next position. The exact definition of a token depends on how the `Scanner` is configured.

By default, the `Scanner` uses white space (spaces, tabs, and newlines) as the delimiter to split the input into tokens. So, when you call `Scan()`, it will read the next token until it encounters a white space character. For example, if you have the input string "Hello World", calling `Scan()` twice would return "Hello" and "World" as separate tokens.

However, you can customize the tokenization behavior by setting a different split function using the `Scanner`'s `Split()` method. You can provide a function that defines your own criteria for what constitutes a token. This allows you to tokenize the input based on specific patterns or delimiters that are relevant to your use case.

In summary, in the context of the `bufio` package, a token represents a meaningful unit of data that is obtained by scanning an input stream using a `Scanner`. It can be a word, a sentence, or any other logical division of data based on the defined tokenization rules.

*/

/*
write code for me in the go programming language to read all of the words from a text file and count the frequency of each word in the text file
*/
// https://chat.openai.com/share/6cb2b004-4cfa-4aad-a2d8-47b0eacd36dd

/*
in the go programming language, I have a map with a string and an int. How do I sort the int values from largest to smallest and see their corresponding string value?
*/
// https://chat.openai.com/share/03a44e91-fc0d-4cdb-884a-c8acd8f440d8

/*

	a := 0
	fmt.Println(a)
	a++
	fmt.Println(a)

	fmt.Println("-------")

	m := make(map[string]int)
	fmt.Println(m["beautiful"])
	m["beautiful"]++
	fmt.Println(m["beautiful"])
	m["beautiful"]++
	fmt.Println(m["beautiful"])
	m["beautiful"]++
	fmt.Println(m["beautiful"])

*/
