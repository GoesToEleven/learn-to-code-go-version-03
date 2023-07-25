// $ ./example2 | cut -c1 | grep '[AB]' | uniq
package main

import (
	"crypto/sha1"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		printHashes("A")
		wg.Done()
	}()

	go func() {
		printHashes("B")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Terminating Program")
}

func printHashes(prefix string) {

	for i := 1; i <= 10; i++ {
		num := strconv.Itoa(i)
		sum := sha1.Sum([]byte(num))
		// 5-digit-number: hex encoded hash
		fmt.Printf("%s: %05d: %x\n", prefix, i, sum)
	}

	fmt.Println("Completed", prefix)
}
