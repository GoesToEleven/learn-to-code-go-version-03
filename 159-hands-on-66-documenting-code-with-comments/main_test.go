package main

import (
	"log"
	"testing"
)

func TestParadise(t *testing.T) {
	got := Paradise("Bali")
	want := "My idea of paradise is Bali"
	if got != want {
		log.Fatalf("error - want %s and got %s", want, got)
	}
}
