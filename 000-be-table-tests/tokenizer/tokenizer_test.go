package tokenizer

import (
	"reflect"
	"testing"
)

// common use case for an anonymous struct
var tokenizeCases = []struct {
	text   string
	tokens []string
}{
	{"Who's on first?", []string{"who", "s", "on", "first"}},
	{"let's go to the beach", []string{"let", "s", "go", "to", "the", "beach"}},
	{"", nil},
}

func TestTokenize(t *testing.T) {
	for _, tc := range tokenizeCases {
		t.Run(tc.text, func(t *testing.T) {
			got := Tokenize(tc.text)
			if !reflect.DeepEqual(got, tc.tokens) {
				t.Fatalf("want %#v, got %#v", tc.tokens, got)
			}

		})
	}
}
