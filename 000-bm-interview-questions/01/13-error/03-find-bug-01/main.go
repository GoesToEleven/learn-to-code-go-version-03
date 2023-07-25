package main

import "log"

type customError struct{}

// Error implements the error interface.
func (c *customError) Error() string {
	return "Find the bug."
}

func app() ([]byte, *customError) {
	return nil, nil
}

func main() {
	var err error
	if _, err = app(); err != nil {
		log.Fatal("Why did this fail?")
	}

	log.Println("No Error")
}
