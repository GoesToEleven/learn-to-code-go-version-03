package main

import "testing"

func TestDoubleInt(t *testing.T) {
	got := doubleInt(7)
	want := 14
	if got != want {
		t.Errorf("got %d & want %d", got, want)
	}
}

// got want
func TestDoubleIntTable(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{3, 6},
		{7, 14},
		{8, 16},
	}

	for _, v := range tests {
		got := doubleInt(v.input)
		want := v.output
		if got != want {
			t.Errorf("got %d & want %d", got, want)
		}
	}
}

// given when should
const succeed = "\u2713"
const failed = "\u2717"

func TestDoubleIntGWS(t *testing.T) {
	t.Log("Given the need to test double int func.")

	input := 4
	got := doubleInt(input)
	want := 8

	t.Logf("\tWhen checking %d doubled equals %d", input, got)

	if got == want {
		t.Logf("\t%s\tDoubleInt worked.", succeed)
	} else {
		t.Logf("\t%s\tDoubleInt got %d and wanted %d.", failed, got, want)
	}
}
