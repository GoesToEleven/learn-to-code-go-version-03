package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	s := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !s.Scan() {
			break
		}
		fmt.Println(t.Format(s.Text()))
	}

}
