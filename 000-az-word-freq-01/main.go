package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

func main() {
	f, err := os.Open("great-gatsby.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	m, err := freq(f)
	if err != nil {
		log.Fatalf("error from freq in main: %s", err)
	}
	fmt.Printf("m\t%#v\n", m)
}

var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

func freq(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	lnum := 0
	for s.Scan() {
		s.Text()
		lnum++
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	fmt.Println("num lines:", lnum)
	return nil, nil
}
