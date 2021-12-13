package main

import (
	"sort"
)

func Dec10A(input []string) int {
	sum := 0

	for _, v := range input {
		hasError, char := getSyntaxError(v)

		if hasError {
			points := get10APoints(char)
			sum += points
		}
	}

	return sum
}

func Dec10B(input []string) int {
	scores := []int{}

	for _, v := range input {
		hasError, _ := getSyntaxError(v)

		if hasError {
			continue
		}

		closingCharacters := getClosingCharacters(v)
		points := getClosingCharactersScore(closingCharacters)

		scores = append(scores, points)
	}

	sort.Ints(scores)
	middleIndex := (len(scores) - 1) / 2

	return scores[middleIndex]
}

func getClosingCharactersScore(str string) int {
	score := 0

	for _, v := range str {
		char := string(v)
		score *= 5

		if char == ")" {
			score++
		}

		if char == "]" {
			score += 2
		}

		if char == "}" {
			score += 3
		}

		if char == ">" {
			score += 4
		}
	}

	return score
}

func getSyntaxError(str string) (bool, string) {
	openingCharacters := []string{}

	for _, v := range str {

		char := string(v)

		if char == "(" || char == "[" || char == "{" || char == "<" {
			openingCharacters = append(openingCharacters, char)

			continue
		}

		lastChar := openingCharacters[len(openingCharacters)-1]

		if char == ")" && lastChar != "(" {
			return true, char
		}

		if char == "]" && lastChar != "[" {
			return true, char
		}

		if char == "}" && lastChar != "{" {
			return true, char
		}

		if char == ">" && lastChar != "<" {
			return true, char
		}

		openingCharacters = openingCharacters[:len(openingCharacters)-1]
	}

	return false, ""
}

func getClosingCharacters(str string) string {
	openingCharacters := ""

	for _, v := range str {
		char := string(v)

		if char == "(" || char == "[" || char == "{" || char == "<" {
			openingCharacters += char
			continue
		}

		openingCharacters = openingCharacters[:len(openingCharacters)-1]
	}

	openingCharacters = reverseString(openingCharacters)

	closingCharacters := ""

	for _, v := range openingCharacters {
		char := string(v)

		if char == "(" {
			closingCharacters += ")"
		}

		if char == "[" {
			closingCharacters += "]"
		}

		if char == "{" {
			closingCharacters += "}"
		}

		if char == "<" {
			closingCharacters += ">"
		}
	}

	return closingCharacters
}

func get10APoints(str string) int {
	if str == ")" {
		return 3
	}

	if str == "]" {
		return 57
	}

	if str == "}" {
		return 1197
	}

	return 25137
}
