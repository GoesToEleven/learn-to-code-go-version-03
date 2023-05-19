package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// curl -i -H 'User-Agent: go' https://api.github.com/users/GoesToEleven

func main() {
	resp, err := http.Get("https://api.github.com/users/GoesToEleven")
	if err != nil {
		log.Fatalf("error getting request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: %s", resp.Status)
	}

	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatalf("error - unable to copy: %s", err)
	}
}
