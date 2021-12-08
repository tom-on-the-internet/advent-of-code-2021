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

		for i := 0; i < len(counts)-1; i++ {
			counts[i] = counts[i+1]
		}

		counts[6] += newFish
		counts[8] = newFish
	}

	return sliceSum(counts)
}
