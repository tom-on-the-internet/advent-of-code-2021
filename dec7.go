package main

func Dec7A(input []int) int {
	costs := getCostsA(input)

	return getMin(costs)
}

func Dec7B(input []int) int {
	costs := getCostsB(input)

	return getMin(costs)
}

func getCostsA(input []int) []int {
	max := getMax(input)
	costs := make([]int, max)

	for i := range costs {
		costs[i] = getSimpleCost(input, i)
	}

	return costs
}

func getCostsB(input []int) []int {
	max := getMax(input)
	costs := make([]int, max)

	for i := range costs {
		costs[i] = getComplexCost(input, i)
	}

	return costs
}

func getSimpleCost(input []int, pos int) int {
	cost := 0

	for _, v := range input {
		cost += abs(v - pos)
	}

	return cost
}

func getComplexCost(input []int, pos int) int {
	cost := 0

	for _, v := range input {
		distance := abs(v - pos)
		cost += getCostFromDistance(distance)
	}

	return cost
}

func getCostFromDistance(distance int) int {
	cost := 0
	costPerStep := 0
	pos := 0

	for pos != distance {
		costPerStep++
		cost += costPerStep
		pos++
	}

	return cost
}

func getMax(input []int) int {
	max := 0

	for _, v := range input {
		if v > max {
			max = v
		}
	}

	return max
}

func getMin(input []int) int {
	min := input[0]

	for _, v := range input {
		if v < min {
			min = v
		}
	}

	return min
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
