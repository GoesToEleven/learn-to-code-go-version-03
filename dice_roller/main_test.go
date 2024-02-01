package main

import "testing"

func TestParseInput(t *testing.T) {
	tests := []struct {
		input                                               string
		expectedNumDice, expectedNumSides, expectedModifier int
		expectErr                                           bool
	}{
		{"3d6+1", 3, 6, 1, false},
		{"2d20-5", 2, 20, -5, false},
		{"5d10", 5, 10, 0, false},
		{"invalid", 0, 0, 0, true},
	}

	for _, test := range tests {
		numDice, numSides, modifier, err := parseInput(test.input)
		if (err != nil) != test.expectErr {
			t.Errorf("parseInput(%q) error = %v, expectErr %v", test.input, err, test.expectErr)
			continue
		}
		if numDice != test.expectedNumDice || numSides != test.expectedNumSides || modifier != test.expectedModifier {
			t.Errorf("parseInput(%q) = %v, %v, %v, want %v, %v, %v", test.input, numDice, numSides, modifier, test.expectedNumDice, test.expectedNumSides, test.expectedModifier)
		}
	}
}

func parseInput(s string) (numDice, numSides, modifier int, err error) {
	// Your code here. Don't forget to return the appropriate values at the end.
	return 0, 0, 0, nil // Placeholder return statement. Replace with your actual implementation.
}
