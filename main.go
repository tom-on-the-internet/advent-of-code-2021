package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

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
