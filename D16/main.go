package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x    int
	y    int
	dirX int
	dirY int
}

func main() {
	fd, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		panic(err)
	}
	defer fd.Close()

	fileReader := bufio.NewScanner(fd)

	m := make([][]rune, 0)
	for fileReader.Scan() {
		line := fileReader.Text()
		m = append(m, []rune(line))
	}

	ansArr := make([]int, 0)
	max := 0

	for i := 0; i < len(m); i++ {
		point := Point{i, 0, 0, 1}
		val := solveForPoint(m, point)
		if val > max {
			max = val
		}
		ansArr = append(ansArr, val)

		point = Point{i, len(m[0]) - 1, 0, -1}
		val = solveForPoint(m, point)
		if val > max {
			max = val
		}
		ansArr = append(ansArr, val)
	}

	for i := 0; i < len(m[0]); i++ {
		point := Point{0, i, 1, 0}
		val := solveForPoint(m, point)
		if val > max {
			max = val
		}
		ansArr = append(ansArr, val)

		point = Point{len(m) - 1, i, -1, 0}
		val = solveForPoint(m, point)
		if val > max {
			max = val
		}
		ansArr = append(ansArr, val)
	}

	fmt.Println("All corners result: ", ansArr)
	fmt.Println("part1: ", ansArr[0])
	fmt.Println("part2: ", max)
}

func solveForPoint(m [][]rune, p Point) int {
	copyM := make([][]rune, len(m))
	for i := 0; i < len(m); i++ {
		copyM[i] = make([]rune, len(m[i]))
		copy(copyM[i], m[i])
	}

	visited := make(map[Point]bool)
	solve(m, copyM, p, visited)
	return findNumberOfEnergized(copyM)
}

func findNumberOfEnergized(m [][]rune) int {
	total := 0
	for _, line := range m {
		for _, r := range line {
			if r == '#' {
				total++
			}
		}
	}
	return total
}

func debug(m [][]rune) {
	for _, line := range m {
		fmt.Println(string(line))
	}
}

func solve(m [][]rune, energizedMap [][]rune, p Point, visited map[Point]bool) {
	for {
		if !isWithinBounds(m, p) {
			return
		}

		energizedMap[p.x][p.y] = '#'
		if m[p.x][p.y] != '.' && !canPassThrough(m[p.x][p.y], p) {
			break
		}
		p.x += p.dirX
		p.y += p.dirY
	}

	if ok, _ := visited[p]; ok {
		return
	}
	points := getDirectionBasedOnMirror(m[p.x][p.y], p)
	visited[p] = true
	for _, point := range points {
		point.x += point.dirX
		point.y += point.dirY
		solve(m, energizedMap, point, visited)
	}
}

func isWithinBounds(m [][]rune, p Point) bool {
	return p.x >= 0 && p.x < len(m) && p.y >= 0 && p.y < len(m[0])
}

func canPassThrough(mirror rune, p Point) bool {
	if mirror == '-' {
		return p.dirY != 0
	} else if mirror == '|' {
		return p.dirX != 0
	}
	return false
}

func getDirectionBasedOnMirror(mirror rune, p Point) []Point {
	points := make([]Point, 0)
	if mirror == '/' || mirror == '\\' {
		delta := 1
		if mirror != '/' {
			delta = -1
		}
		if p.dirX == 1 {
			points = append(points, Point{p.x, p.y, 0, -1 * delta})
		} else if p.dirX == -1 {
			points = append(points, Point{p.x, p.y, 0, 1 * delta})
		} else if p.dirY == 1 {
			points = append(points, Point{p.x, p.y, -1 * delta, 0})
		} else if p.dirY == -1 {
			points = append(points, Point{p.x, p.y, 1 * delta, 0})
		}
	} else if mirror == '-' && p.dirY == 0 {
		points = append(points, Point{p.x, p.y, 0, 1})
		points = append(points, Point{p.x, p.y, 0, -1})
	} else if mirror == '|' && p.dirX == 0 {
		points = append(points, Point{p.x, p.y, 1, 0})
		points = append(points, Point{p.x, p.y, -1, 0})
	}
	return points
}
