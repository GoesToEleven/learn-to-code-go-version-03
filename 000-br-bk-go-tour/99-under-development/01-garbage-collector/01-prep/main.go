package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	numCPU := runtime.NumCPU()
	fmt.Printf("Number of CPUs: %d\n", numCPU)

	t := time.Now()
	fmt.Println(t)

	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if ! s.Scan() {
			break
		}
		fmt.Println(t.Format(s.Text()))
	}
	
}