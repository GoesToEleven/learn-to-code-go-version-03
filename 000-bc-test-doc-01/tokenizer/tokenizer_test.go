package tokenizer

import (
	"reflect"
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
