package main

import (
	"strings"
)

type (
	grid  map[point]int
	point struct {
		x int
		y int
	}
)

type line struct {
	start point
	end   point
}

func (l line) isStraight() bool {
	return l.start.x == l.end.x || l.start.y == l.end.y
}

func (l line) points() []point {
	points := []point{}

	x := l.start.x
	y := l.start.y

	for x != l.end.x || y != l.end.y {
		points = append(points, point{x: x, y: y})

		if x != l.end.x && l.start.x < l.end.x {
			x++
		}

		if x != l.end.x && l.start.x > l.end.x {
			x--
		}

		if y != l.end.y && l.start.y < l.end.y {
			y++
		}

		if y != l.end.y && l.start.y > l.end.y {
			y--
		}
	}

	points = append(points, point{x: x, y: y})

	return points
}

func Dec5A(input []string) int {
	lines := makeLines(input)
	straightLines := getStraight(lines)

	grid := makeGrid(straightLines)

	return countOverlaps(grid)
}

func Dec5B(input []string) int {
	lines := makeLines(input)

	grid := makeGrid(lines)

	return countOverlaps(grid)
}

func makeLines(input []string) []line {
	lines := []line{}

	for _, v := range input {
		ps := strings.Split(v, " -> ")

		line := line{
			start: strsToPoint(ps[0]),
			end:   strsToPoint(ps[1]),
		}

		lines = append(lines, line)
	}

	return lines
}

func strsToPoint(strs string) point {
	q := strings.Split(strs, ",")

	return point{x: strToInt(q[0]), y: strToInt(q[1])}
}

func getStraight(lines []line) []line {
	straight := []line{}

	for _, v := range lines {
		if v.isStraight() {
			straight = append(straight, v)
		}
	}

	return straight
}

func makeGrid(lines []line) grid {
	g := grid{}

	for _, line := range lines {
		for _, point := range line.points() {
			_, ok := g[point]
			if !ok {
				g[point] = 0
			}

			g[point]++
		}
	}

	return g
}

func countOverlaps(g grid) int {
	count := 0

	for _, v := range g {
		if v > 1 {
			count++
		}
	}

	return count
}
