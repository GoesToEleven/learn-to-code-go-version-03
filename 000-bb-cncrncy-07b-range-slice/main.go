package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	// siteTime("https://google.com")

	xs := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.apple.com",
		"https://www.reddit.com",
	}

	ch := make(chan string)
	for _, v := range xs {
		go siteTime(v, ch)
	}

	// range loop over a slice
	// it will loop the same number of times as the items that need to be pulled from the channel
	for range xs {
		fmt.Println(<-ch)
	}

}

func siteTime(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("ERROR: %s --> %s", url, err)
		return
	}
	defer resp.Body.Close()

	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		log.Printf("ERROR: %s --> %s", url, err)
	}

	elapsed := time.Since(start)
	ch <- fmt.Sprintf("INFO: %s --> %v", url, elapsed)
}
