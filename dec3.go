package main

import (
	"strconv"
)

func Dec3A(strs []string) int {
	gamma := ""
	epsilon := ""

	for i := 0; i <= len(strs[0])-1; i++ {
		zeros := 0
		ones := 0

		for _, v := range strs {
			if string(v[i]) == "0" {
				zeros++
			} else {
				ones++
			}
		}

		if zeros > ones {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	decGamma, _ := strconv.ParseInt(gamma, 2, 64)
	decEpsilon, _ := strconv.ParseInt(epsilon, 2, 64)

	return int(decGamma * decEpsilon)
}

// This solution loops 4 times
// And is generally messy.
func Dec3B(strs []string) int {
	return calculateCSR(strs) * calculateOGR(strs)
}

func calculateOGR(strs []string) int {
	position := 0

	for len(strs) != 1 {
		strs = getRemainingCommon(strs, position, getMostCommon(strs, position))
		position++
	}

	num, _ := strconv.ParseInt(strs[0], 2, 64)

	return int(num)
}

func calculateCSR(strs []string) int {
	position := 0

	for len(strs) != 1 {
		strs = getRemainingNotCommon(strs, position, getMostCommon(strs, position))
		position++
	}

	num, _ := strconv.ParseInt(strs[0], 2, 64)

	return int(num)
}

func getMostCommon(strs []string, index int) string {
	sum := 0

	for _, v := range strs {
		if string(v[index]) == "0" {
			sum--
		} else {
			sum++
		}
	}

	if sum < 0 {
		return "0"
	}

	return "1"
}

func getRemainingCommon(strs []string, index int, common string) []string {
	remaining := []string{}

	for _, v := range strs {
		if string(v[index]) == common {
			remaining = append(remaining, v)
		}
	}

	return remaining
}

func getRemainingNotCommon(strs []string, index int, common string) []string {
	remaining := []string{}

	for _, v := range strs {
		if string(v[index]) != common {
			remaining = append(remaining, v)
		}
	}

	return remaining
}
