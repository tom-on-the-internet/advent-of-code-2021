package main

// Dec1A calculates the number of times the depth
// increased from it's previous depth.
func Dec1A(depths []int) int {
	count := 0

	for i, v := range depths {
		if i == 0 {
			continue
		}

		if v > depths[i-1] {
			count++
		}
	}

	return count
}

// Dec1B calculates the number of times the depth
// increased from it's previous depth use sliding window.
func Dec1B(depths []int) int {
	count := 0

	for i, v := range depths {
		if i < 3 {
			continue
		}

		if v > depths[i-3] {
			count++
		}
	}

	return count
}
