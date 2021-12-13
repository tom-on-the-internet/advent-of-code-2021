package main

import "testing"

func TestDec9A(t *testing.T) {
	tests := map[string]struct {
		got  []string
		want int
	}{
		"simple": {got: []string{
			"2199943210",
			"3987894921",
			"9856789892",
			"8767896789",
			"9899965678",
		}, want: 15},
		"from file": {got: fileToStringSlice("dec9.txt"), want: 530},
	}

	for name, tc := range tests {
		got := Dec9A(tc.got)
		if got != tc.want {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}

func TestDec9B(t *testing.T) {
	tests := map[string]struct {
		got  []string
		want int
	}{
		"simple": {got: []string{
			"2199943210",
			"3987894921",
			"9856789892",
			"8767896789",
			"9899965678",
		}, want: 1134},
		"from file": {got: fileToStringSlice("dec9.txt"), want: 1019494},
	}

	for name, tc := range tests {
		got := Dec9B(tc.got)
		if got != tc.want {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}
