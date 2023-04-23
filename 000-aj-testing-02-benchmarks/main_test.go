package main

import "testing"

func TestSum(t *testing.T) {
	got := sum(2, 3)
	want := 5
	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}

func TestMultiply(t *testing.T) {
	got := multiply(2, 3)
	want := 6
	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}

// test table
func TestSumTable(t *testing.T) {
	tests := []struct {
		xi   []int
		want int
	}{
		{[]int{2, 3}, 5},
		{[]int{4, 6}, 10},
		{[]int{11, 12}, 23},
	}

	for _, v := range tests {
		got := sum(v.xi[0], v.xi[1])
		want := v.want
		if got != want {
			t.Errorf("got %v wanted %v", got, want)
		}
	}
}

func BenchmarkMassAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		massAdd()
	}
}

/* https://pkg.go.dev/testing#hdr-Benchmarks
benchmarks are executed by the "go test" command
when its -bench flag is provided.

to run the benchmark test
enter this at the command line

go test -bench .

for help, enter this at the command line
go help
go help testflag

*/
