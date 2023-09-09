package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	var y int
	fmt.Println(y)
	
	x := 42
	fmt.Println(x)

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

	var zz int
	fmt.Println(zz)
}
