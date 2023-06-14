package main

import (
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	// siteTime("https://google.com")

	g := new(errgroup.Group)
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.apple.com",
		"https://www.reddit.com",
		"https://www.doesnotexistsite.biz",
	}

	for _, url := range urls {
		url := url
		g.Go(func() error {
			log.Printf("GETTING URL: %s", url)
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}

	if err := g.Wait(); err == nil {
		log.Printf("Successfully fetched all URLs. %s \t %#v\n", err, err)
	} else {
		log.Printf("There was an error: %s \t %#v\n", err, err)
	}
}
