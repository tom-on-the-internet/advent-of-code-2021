package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type matrix [][]int

func (m matrix) xlen() int {
	return len(m[0])
}

func (m matrix) ylen() int {
	return len(m)
}

func main() {
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

func fileToIntSlice(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numbers := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}

	return numbers
}

func strToInt(s string) int {
	x, _ := strconv.Atoi(s)

	return x
}

func csvToInts(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numbers := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strNums := strings.Split(scanner.Text(), ",")

		for _, v := range strNums {
			numbers = append(numbers, strToInt(v))
		}
	}

	return numbers
}

func sliceSum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	return sum
}

func toMatrix(input []string) matrix {
	matrix := matrix{}

	for _, v := range input {
		split := strings.Split(v, "")
		ints := []int{}

		for _, w := range split {
			i, _ := strconv.Atoi(w)
			ints = append(ints, i)
		}

		matrix = append(matrix, ints)
	}

	return matrix
}
