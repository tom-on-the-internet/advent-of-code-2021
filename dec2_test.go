package main

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestDec2A(t *testing.T) {
	tests := map[string]struct {
		got  []string
		want int
	}{
		"simple": {got: []string{
			"forward 5",
			"down 5",
			"forward 8",
			"up 3",
			"down 8",
			"forward 2",
		}, want: 150},
		"from file": {got: fileToStringSlice("dec2.txt"), want: 1488669},
	}

	for name, tc := range tests {
		got := Dec2A(tc.got)
		if got != tc.want {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}

func TestDec2B(t *testing.T) {
	tests := map[string]struct {
		got  []string
		want int
	}{
		"simple": {got: []string{
			"forward 5",
			"down 5",
			"forward 8",
			"up 3",
			"down 8",
			"forward 2",
		}, want: 900},
		"from file": {got: fileToStringSlice("dec2.txt"), want: 1176514794},
	}

	for name, tc := range tests {
		got := Dec2B(tc.got)
		if got != tc.want {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}

func fileToStringSlice(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	strings := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	return strings
}
