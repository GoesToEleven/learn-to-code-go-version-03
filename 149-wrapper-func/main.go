package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	xb, err := readFile("poem.txt")
	if err != nil {
		log.Fatalf("error in main in readFile: %s", err)
	}
	fmt.Println(xb)
	fmt.Println(string(xb))
}

func readFile(fileName string) ([]byte, error) {
	xb, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("there was an error in readfile: %s", err)
	}
	return xb, nil
}
