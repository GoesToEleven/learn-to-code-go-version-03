package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func _()

func main() {
	// routing
	// /health is an exact match
	// /health/ is a prefix match
	http.HandleFunc("/hi", greet)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

// go run main.go

// open a web browser, go to http://localhost:8080/hi

// at another terminal
// curl http://localhost:8080

// at a terminal
// go install github.com/rakyll/hey@latest
// hey -n 10000 http://localhost:8080/hi