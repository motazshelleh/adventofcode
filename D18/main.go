package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	dir    rune
	length int
	color  string
}

func main() {
	fd, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		panic(err)
	}
	defer fd.Close()

	fileReader := bufio.NewScanner(fd)

	points := make([]Point, 0)
	for fileReader.Scan() {
		line := fileReader.Text()
		parts := strings.Split(line, " ")
		length, _ := strconv.Atoi(parts[1])
		points = append(points, Point{dir: rune(parts[0][0]), length: length, color: parts[2]})
	}

	for _, point := range points {
		fmt.Println(point)
	}

	maxX, maxY, startX, startY := findMapSize(points)
	fmt.Println("Map size:", maxX, maxY)

	m := drawMap(points, maxX, maxY, startX, startY)

	for i, row := range m {
		for j, val := range row {
			if val == "" {
				isClosed := floodFill(m, i, j)
				replacement := "."
				if isClosed {
					replacement = "#"
				}
				replacePlaceholders(m, replacement)
			}
		}
	}

	// debugMap(m)

	count := 0
	for _, row := range m {
		for _, val := range row {
			if val == "#" {
				count++
			}
		}
	}
	fmt.Println("Count:", count)
}

func replacePlaceholders(m [][]string, replacement string) {
	for i, row := range m {
		for j, val := range row {
			if val == "O" {
				m[i][j] = replacement
			}
		}
	}
}

func drawMap(points []Point, y, x int, startX int, startY int) [][]string {
	// Create map
	m := make([][]string, x)
	for i := range m {
		m[i] = make([]string, y)
	}

	ptrX, ptrY := startX, startY
	for _, point := range points {
		deltaX, deltaY := 0, 0
		if point.dir == 'R' {
			deltaY = 1
		} else if point.dir == 'L' {
			deltaY = -1
		} else if point.dir == 'U' {
			deltaX = -1
		} else if point.dir == 'D' {
			deltaX = 1
		}
		for i := 0; i < point.length; i++ {
			ptrX += deltaX
			ptrY += deltaY
			m[ptrX][ptrY] = "#"
		}
	}

	return m
}

func floodFill(m [][]string, x, y int) bool {
	if x < 0 || x >= len(m) || y < 0 || y >= len(m[0]) {
		return false
	}
	if m[x][y] != "#" && m[x][y] != "O" {
		m[x][y] = "O"
		return floodFill(m, x+1, y) && floodFill(m, x-1, y) && floodFill(m, x, y+1) && floodFill(m, x, y-1)
	}
	return true
}

func debugMap(m [][]string) {
	for _, row := range m {
		for _, val := range row {
			if val == "" {
				fmt.Print(".")
			} else {
				fmt.Print(val)
			}
		}
		fmt.Println()
	}
}

func findMapSize(points []Point) (int, int, int, int) {
	ptrX, ptrY := 0, 0
	maxX, maxY := 0, 0
	minX, minY := 0, 0

	for _, point := range points {
		deltaX, deltaY := 0, 0
		if point.dir == 'R' {
			deltaY = 1
		} else if point.dir == 'L' {
			deltaY = -1
		} else if point.dir == 'U' {
			deltaX = -1
		} else if point.dir == 'D' {
			deltaX = 1
		}
		for i := 0; i < point.length; i++ {
			ptrX += deltaX
			ptrY += deltaY
		}

		if ptrX > maxX {
			maxX = ptrX
		}
		if ptrY > maxY {
			maxY = ptrY
		}
		if ptrX < minX {
			minX = ptrX
		}
		if ptrY < minY {
			minY = ptrY
		}
	}

	fmt.Println("MinX:", minX, "MaxX:", maxX, "MinY:", minY, "MaxY:", maxY)
	return maxY - minY + 1, maxX - minX + 1, minX * -1, minY * -1
}
