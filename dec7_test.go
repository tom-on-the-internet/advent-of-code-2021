package main

import "testing"

func TestDec7A(t *testing.T) {
	tests := map[string]struct {
		got  []int
		want int
	}{
		"simple": {got: []int{
			16, 1, 2, 0, 4, 2, 7, 1, 2, 14,
		}, want: 37},
		"from file": {got: csvToInts("dec7.txt"), want: 343468},
	}

	for name, tc := range tests {
		got := Dec7A(tc.got)
		if got != tc.want {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}

func TestDec7B(t *testing.T) {
	tests := map[string]struct {
		got  []int
		want int
	}{
		"simple": {got: []int{
			16, 1, 2, 0, 4, 2, 7, 1, 2, 14,
		}, want: 168},
		"from file": {got: csvToInts("dec7.txt"), want: 96086265},
	}

	for name, tc := range tests {
		got := Dec7B(tc.got)
		if got != tc.want {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}
