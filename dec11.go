package main

func Dec11A(input []string) int {
	m := toMatrix(input)
	sum := 0

	for i := 1; i <= 100; i++ {
		for _, v := range m {
			for j := range v {
				v[j]++
			}
		}

		flashes := 0
		m, flashes = calcFlashes(m, flashes)
		sum += flashes
	}

	return sum
}

func Dec11B(input []string) int {
	m := toMatrix(input)
	step := 0

	for {
		step++
		for _, v := range m {
			for j := range v {
				v[j]++
			}
		}

		m, _ = calcFlashes(m, 0)

		if allFlashed(m) {
			return step
		}

	}

	return 0
}

func allFlashed(m matrix) bool {
	for _, v := range m {
		for _, w := range v {
			if w != 0 {
				return false
			}
		}
	}

	return true
}

func calcFlashes(m matrix, flashes int) (matrix, int) {
	if !readyToFlash(m) {
		return m, flashes
	}

	for i, v := range m {
		for j, w := range v {
			if w > 9 {
				flashes++

				v[j] = 0

				// TL
				if i > 0 && j > 0 {
					if m[i-1][j-1] != 0 {
						m[i-1][j-1]++
					}
				}

				// T
				if i > 0 {
					if m[i-1][j] != 0 {
						m[i-1][j]++
					}
				}

				// TR
				if i > 0 && j < len(v)-1 {
					if m[i-1][j+1] != 0 {
						m[i-1][j+1]++
					}
				}

				// ML
				if j > 0 {
					if m[i][j-1] != 0 {
						m[i][j-1]++
					}
				}

				// MR
				if j < len(v)-1 {
					if m[i][j+1] != 0 {
						m[i][j+1]++
					}
				}

				// BL
				if i < len(m)-1 && j > 0 {
					if m[i+1][j-1] != 0 {
						m[i+1][j-1]++
					}
				}

				// B
				if i < len(m)-1 {
					if m[i+1][j] != 0 {
						m[i+1][j]++
					}
				}

				// BR
				if i < len(m)-1 && j < len(v)-1 {
					if m[i+1][j+1] != 0 {
						m[i+1][j+1]++
					}
				}
			}
		}
	}

	return calcFlashes(m, flashes)
}

func readyToFlash(m matrix) bool {
	for _, v := range m {
		for _, w := range v {
			if w > 9 {
				return true
			}
		}
	}

	return false
}
