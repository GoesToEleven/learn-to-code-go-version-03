package main

import (
	"io"
	"log"
	"net/http"
	"sync"
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

	var wg sync.WaitGroup
	wg.Add(len(xs))
	for _, v := range xs {
		v:=v
		go func(){
			defer wg.Done()
			siteTime(v)
		}()
	}
	wg.Wait()
}

func siteTime(url string) {
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
	log.Printf("INFO: %s --> %v", url, elapsed)
}
