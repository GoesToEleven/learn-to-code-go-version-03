package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestRollDice(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"2d6", 7},     // Simple roll
		{"1d20+5", 18}, // Roll with modifier
		{"4d4-2", 9},   // Multiple dice with modifier
	}

	for _, test := range tests {
		// Capture standard output to compare results
		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Call the main function with the test input
		main()

		// Restore standard output
		w.Close()
		os.Stdout = old

		// Read the captured output
		output, _ := ioutil.ReadAll(r)
		resultLine := strings.Split(string(output), "\n")[len(strings.Split(string(output), "\n"))-1]
		splitResult := strings.Split(resultLine, ": ")
		if len(splitResult) >= 2 {
			actual, _ := strconv.Atoi(splitResult[1])
			// Use actual here
			if actual != test.expected {
				t.Errorf("Roll failed for input %s: expected %d, got %d", test.input, test.expected, actual)
			}
		}
	}
}
