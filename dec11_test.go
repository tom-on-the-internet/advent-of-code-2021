package main

import "testing"

func TestDec11A(t *testing.T) {
	tests := map[string]struct {
		got  []string
		want int
	}{
		"simple": {got: []string{
			"5483143223",
			"2745854711",
			"5264556173",
			"6141336146",
			"6357385478",
			"4167524645",
			"2176841721",
			"6882881134",
			"4846848554",
			"5283751526",
		}, want: 1656},
		"from file": {got: fileToStringSlice("dec11.txt"), want: 1617},
	}

	for name, tc := range tests {
		got := Dec11A(tc.got)
		if got != tc.want {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}

func TestDec11B(t *testing.T) {
	tests := map[string]struct {
		got  []string
		want int
	}{
		"simple": {got: []string{
			"5483143223",
			"2745854711",
			"5264556173",
			"6141336146",
			"6357385478",
			"4167524645",
			"2176841721",
			"6882881134",
			"4846848554",
			"5283751526",
		}, want: 195},
		"from file": {got: fileToStringSlice("dec11.txt"), want: 258},
	}

	for name, tc := range tests {
		got := Dec11B(tc.got)
		if got != tc.want {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}
