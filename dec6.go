package main

func Dec6A(input []int) int {
	days := 80

	return breed(input, days)
}

func Dec6B(input []int) int {
	days := 256

	return breed(input, days)
}

func breed(input []int, days int) int {
	counts := make([]int, 9)
	for _, v := range input {
		counts[v]++
	}

	for x := 1; x <= days; x++ {
		newFish := counts[0]

		counts[0] = counts[1]
		counts[1] = counts[2]
		counts[2] = counts[3]
		counts[3] = counts[4]
		counts[4] = counts[5]
		counts[5] = counts[6]
		counts[6] = counts[7] + newFish
		counts[7] = counts[8]
		counts[8] = newFish
	}

	return sliceSum(counts)
}
