package main

import (
	"testing"
)

func TestDec3A(t *testing.T) {
	tests := map[string]struct {
		got  []string
		want int
	}{
		"simple": {got: []string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		}, want: 198},
		"from file": {got: fileToStringSlice("dec3.txt"), want: 1540244},
	}

	for name, tc := range tests {
		got := Dec3A(tc.got)
		if got != tc.want {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}

func TestDec3B(t *testing.T) {
	tests := map[string]struct {
		got  []string
		want int
	}{
		"simple": {got: []string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		}, want: 230},
		"from file": {got: fileToStringSlice("dec3.txt"), want: 4203981},
	}

	for name, tc := range tests {
		got := Dec3B(tc.got)
		if got != tc.want {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}
