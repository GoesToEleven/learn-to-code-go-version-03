package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// new file
	f, err := os.Create("temp.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//write to new file
	_, err = f.Write([]byte("this text goes in the file!"))
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println("done")
}
