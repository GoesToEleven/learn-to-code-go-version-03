package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"mymodule/000-bj-http-02-tokenizer/tokenizer"
)

func _()

func main() {
	// routing
	// /health is an exact match
	// /health/ is a prefix match
	http.HandleFunc("/hi", greet)
	http.HandleFunc("/tokenize", tokenizeHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

	//step 1: get, convert, & validate data
	defer r.Body.Close()

	// good idea to use
	// rdr := io.LimitReader(r.Body, 1_000_000)

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "can't read", http.StatusBadRequest)
		return
	}

	text := string(data)

	//step 2: work
	tokens := tokenizer.Tokenize(text)

	//step 3: encode & emit output
	resp := map[string]any{
		"tokens": tokens,
	}

	// you can also do
	// err = json.NewEncoder(w).Encode(resp)

	data, err = json.Marshal(resp)
	if err != nil {
		http.Error(w, "can't encode", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}

// go run main.go

// at another terminal
// curl -i -d"Let's go to the beach" http://localhost:8080/tokenize

// at a terminal
// go install github.com/rakyll/hey@latest
// hey -n 10000 http://localhost:8080/hi
