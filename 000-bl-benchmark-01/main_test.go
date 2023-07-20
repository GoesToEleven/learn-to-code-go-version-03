package main

import "testing"

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(42, 13)
	}
}
