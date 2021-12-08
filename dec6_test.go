package main

import "testing"

func TestDec6A(t *testing.T) {
	tests := map[string]struct {
		got  []int
		want int
	}{
		"simple": {got: []int{
			3, 4, 3, 1, 2,
		}, want: 5934},
		"from file": {got: csvToInts("dec6.txt"), want: 352151},
	}

	for name, tc := range tests {
		got := Dec6A(tc.got)
		if got != tc.want {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}

func TestDec6B(t *testing.T) {
	tests := map[string]struct {
		got  []int
		want int
	}{
		"simple": {got: []int{
			3, 4, 3, 1, 2,
		}, want: 26984457539},
		"from file": {got: csvToInts("dec6.txt"), want: 1601616884019},
	}

	for name, tc := range tests {
		got := Dec6B(tc.got)
		if got != tc.want {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}
