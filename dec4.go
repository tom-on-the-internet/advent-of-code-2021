package main

import (
	"strconv"
	"strings"
)

func Dec4A(input []string) int {
	numbers := getNumbers(input)

	bingoCards := getBingoCards(input)

	winningCard, lastNumber := findWinningCard(bingoCards, numbers)

	return getValue(winningCard, lastNumber)
}

func Dec4B(input []string) int {
	numbers := getNumbers(input)

	bingoCards := getBingoCards(input)

	losingCard, lastNumber := findLosingCard(bingoCards, numbers)

	return getValue(losingCard, lastNumber)
}

func getValue(card [][]int, number int) int {
	value := 0

	for _, row := range card {
		for _, col := range row {
			if col != -1 {
				value += col
			}
		}
	}

	return value * number
}

func findWinningCard(cards [][][]int, numbers []int) ([][]int, int) {
	winningCard := [][]int{}
	winningNumber := 0

	for _, v := range numbers {
		for i, card := range cards {
			updatedCard, hadNumber := markCard(card, v)
			if hadNumber && cardWon(updatedCard) {
				return updatedCard, v
			}

			cards[i] = updatedCard
		}
	}

	return winningCard, winningNumber
}

func findLosingCard(cards [][][]int, numbers []int) ([][]int, int) {
	losingCard := [][]int{}
	lastNumber := 0

	countWinningCards := 0
	m := make(map[int]bool)

	for _, v := range numbers {
		for i, card := range cards {
			// skip winners
			val, ok := m[i]
			if val && ok {
				continue
			}

			updatedCard, hadNumber := markCard(card, v)
			cards[i] = updatedCard

			if hadNumber && cardWon(updatedCard) {
				m[i] = true

				countWinningCards++
				if countWinningCards == len(cards) {
					return updatedCard, v
				}
			}
		}
	}

	return losingCard, lastNumber
}

func cardWon(card [][]int) bool {
	for i := 0; i <= 4; i++ {
		if card[i][0] == -1 && card[i][1] == -1 && card[i][2] == -1 && card[i][3] == -1 && card[i][4] == -1 {
			return true
		}

		if card[0][i] == -1 && card[1][i] == -1 && card[2][i] == -1 && card[3][i] == -1 && card[4][i] == -1 {
			return true
		}
	}

	return false
}

func getNumbers(input []string) []int {
	str := input[0]
	nums := []int{}

	for _, v := range strings.Split(str, ",") {
		num, _ := strconv.Atoi(strings.TrimSpace(v))
		nums = append(nums, num)
	}

	return nums
}

func getBingoCards(input []string) [][][]int {
	bingoCards := [][][]int{}
	bingoCard := [][]int{}
	input = input[2:]

	for _, row := range input {
		if len(row) == 0 {
			bingoCards = append(bingoCards, bingoCard)
			bingoCard = [][]int{}

			continue
		}

		nums := []int{}

		for _, v := range strings.Fields(row) {
			num, _ := strconv.Atoi(strings.TrimSpace(v))
			nums = append(nums, num)
		}

		bingoCard = append(bingoCard, nums)
	}

	bingoCards = append(bingoCards, bingoCard)

	return bingoCards
}

func markCard(card [][]int, num int) ([][]int, bool) {
	newCard := card
	found := false
	for i, row := range card {
		for j, col := range row {
			if col == num {
				newCard[i][j] = -1
				found = true
			} else {
				newCard[i][j] = col
			}
		}
	}
	return newCard, found
}

func remove(s [][][]int, i int) [][][]int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
