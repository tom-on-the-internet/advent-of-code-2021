package main

import (
	"testing"
)

func TestDec1A(t *testing.T) {
	tests := map[string]struct {
		got  []int
		want int
	}{
		"simple":    {got: []int{1, 3, 4, 2, 5}, want: 3},
		"from file": {got: fileToIntSlice("dec1a.txt"), want: 1759},
	}

	for name, tc := range tests {
		got := Dec1A(tc.got)
		if got != tc.want {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}

func TestDec1B(t *testing.T) {
	tests := map[string]struct {
		got  []int
		want int
	}{
		"simple": {got: []int{
			199,
			200,
			208,
			210,
			200,
			207,
			240,
			269,
			260,
			263,
		}, want: 5},
		"from file": {got: fileToIntSlice("dec1a.txt"), want: 1805},
	}

	for name, tc := range tests {
		got := Dec1B(tc.got)
		if got != tc.want {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}
