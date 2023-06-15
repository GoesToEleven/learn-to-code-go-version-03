package tokenizer

import (
	"reflect"
	"strings"
	"testing"
)

func TestTokenize(t *testing.T) {
	text := "let's go to the beach"
	want := []string{"let", "s", "go", "to", "the", "beach"}
	got := Tokenize(text)
	// if want != got { // can't compare slices to each other; only to nil
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, got %#v", want, got)
	}
}

func FuzzTokenize(f *testing.F) {
	f.Fuzz(func(t *testing.T, text string) {
		tokens := Tokenize(text)
		lText := strings.ToLower(text)
		for _, tok := range tokens {
			if !strings.Contains(lText, tok) {
				t.Fatal(tok)
			}
		}
	})
}

// go test -fuzz . -fuzztime 10s -v
