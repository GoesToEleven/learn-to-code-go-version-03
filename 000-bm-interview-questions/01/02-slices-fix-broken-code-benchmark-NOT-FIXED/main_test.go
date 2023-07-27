package main

import "testing"

func BenchmarkValSemantics(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qty := 10
		vs := make([]string, 0, qty)
		valSemantics(vs, qty)
	}
}

func BenchmarkPtrSemantics(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qty := 10
		ps := make([]string, qty)
		ptrSemantics(ps, qty)
	}

}
