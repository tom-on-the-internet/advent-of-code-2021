package main

import (
	"strconv"
	"strings"
)

func Dec2A(strs []string) int {
	depth := 0
	horizontalPosition := 0

	for _, s := range strs {
		if strings.HasPrefix(s, "forward") {
			v, _ := strconv.Atoi(strings.Split(s, " ")[1])
			horizontalPosition += v

			continue
		}

		if strings.HasPrefix(s, "up") {
			v, _ := strconv.Atoi(strings.Split(s, " ")[1])
			depth -= v

			continue
		}

		if strings.HasPrefix(s, "down") {
			v, _ := strconv.Atoi(strings.Split(s, " ")[1])
			depth += v

			continue
		}
	}

	return depth * horizontalPosition
}

func Dec2B(strs []string) int {
	depth := 0
	horizontalPosition := 0
	aim := 0

	for _, s := range strs {
		if strings.HasPrefix(s, "forward") {
			v, _ := strconv.Atoi(strings.Split(s, " ")[1])
			horizontalPosition += v
			depth += aim * v

			continue
		}

		if strings.HasPrefix(s, "up") {
			v, _ := strconv.Atoi(strings.Split(s, " ")[1])
			aim -= v

			continue
		}

		if strings.HasPrefix(s, "down") {
			v, _ := strconv.Atoi(strings.Split(s, " ")[1])
			aim += v

			continue
		}
	}

	return depth * horizontalPosition
}
