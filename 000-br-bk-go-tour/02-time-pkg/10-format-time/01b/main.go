package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scn := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("> ")
		if !scn.Scan() {
			break
		}
		log.Printf("%s", scn.Text())
	}
}
