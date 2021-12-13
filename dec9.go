package main

import "sort"

func Dec9A(input []string) int {
	m := toMatrix(input)
	points := lowestPoints(m)
	values := lowestPointValues(points, m)

	return sumRiskLevel(values)
}

func Dec9B(input []string) int {
	m := toMatrix(input)
	points := lowestPoints(m)

	sizes := []int{}
	for _, p := range points {
		sizes = append(sizes, basinSize(p, m))
	}

	sort.Ints(sizes)

	l := len(sizes)

	return sizes[l-1] * sizes[l-2] * sizes[l-3]
}

func basinSize(p point, m matrix) int {
	start := p
	seen := map[point]bool{
		start: true,
	}

	queue := []point{start}

	var current point

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		for _, neighbor := range neighbors(current, m) {
			if _, ok := seen[neighbor]; !ok {
				queue = append(queue, neighbor)
				seen[neighbor] = true
			}
		}
	}

	return len(seen)
}

func neighbors(p point, m matrix) []point {
	var result []point

	directions := []point{
		{x: -1, y: 0},
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: 0, y: 1},
	}

	for _, direction := range directions {
		neighborPoint := point{x: p.x + direction.x, y: p.y + direction.y}

		if neighborPoint.x < 0 || neighborPoint.x > m.xlen()-1 || neighborPoint.y < 0 || neighborPoint.y > m.ylen()-1 {
			continue
		}

		// we know the point is in the matrix 😎
		if m[neighborPoint.y][neighborPoint.x] >= m[p.y][p.x] && m[neighborPoint.y][neighborPoint.x] != 9 {
			result = append(result, neighborPoint)
		}
	}

	return result
}

func lowestPoints(m matrix) []point {
	points := []point{}

	for i, v := range m {
		for j, w := range v {
			// fmt.Println(i, j, ":", w)
			// if lower than above
			if i > 0 && w >= m[i-1][j] {
				continue
			}
			// if lower than below
			if i < m.ylen()-1 && w >= m[i+1][j] {
				continue
			}
			// if lower than left
			if j > 0 && w >= m[i][j-1] {
				continue
			}
			// if lower than right
			if j < m.xlen()-1 && w >= m[i][j+1] {
				continue
			}

			points = append(points, point{x: j, y: i})
		}
	}

	return points
}

func lowestPointValues(points []point, m matrix) []int {
	values := []int{}

	for _, p := range points {
		values = append(values, m[p.y][p.x])
	}

	return values
}

func sumRiskLevel(points []int) int {
	amt := 0

	for _, v := range points {
		amt += v + 1
	}

	return amt
}
