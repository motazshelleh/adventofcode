package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func main() {
	fd, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		panic(err)
	}
	defer fd.Close()

	fileReader := bufio.NewScanner(fd)

	inputMap := make([][]rune, 0)
	inputMapExpanded := make([][]rune, 0)
	for fileReader.Scan() {
		line := fileReader.Text()
		inputMap = append(inputMap, []rune(line))
		inputMapExpanded = append(inputMapExpanded, make([]rune, 0))
		isAllDots := true
		for i := 0; i < len(line); i++ {
			if line[i] != '.' {
				isAllDots = false
				break
			}
		}
		if isAllDots {
			inputMap = append(inputMap, []rune(line))
			inputMapExpanded = append(inputMapExpanded, make([]rune, 0))
		}
	}
	// for i := 0; i < len(inputMap); i++ {
	// 	fmt.Println(string(inputMap[i]))
	// }

	for i := 0; i < len(inputMap[0]); i++ {
		isAllDots := true
		for j := 0; j < len(inputMap); j++ {
			if inputMap[j][i] != '.' {
				isAllDots = false
			}
			inputMapExpanded[j] = append(inputMapExpanded[j], inputMap[j][i])
		}
		if isAllDots {
			for j := 0; j < len(inputMap); j++ {
				inputMapExpanded[j] = append(inputMapExpanded[j], inputMap[j][i])
			}
		}

	}

	// for i := 0; i < len(inputMapExpanded); i++ {
	// 	fmt.Println(string(inputMapExpanded[i]))
	// }

	glaxiesPosition := make([]Point, 0)
	for i := 0; i < len(inputMapExpanded); i++ {
		for j := 0; j < len(inputMapExpanded[i]); j++ {
			if inputMapExpanded[i][j] != '.' {
				p := Point{j, i}
				glaxiesPosition = append(glaxiesPosition, p)
			}
		}
	}

	// fmt.Println(glaxiesPosition)

	sumOfManhattanDistance := 0
	for i := 0; i < len(glaxiesPosition); i++ {
		for j := i; j < len(glaxiesPosition); j++ {
			sumOfManhattanDistance += manhattanDistance(glaxiesPosition[i], glaxiesPosition[j])
		}
	}

	fmt.Println(sumOfManhattanDistance)
}

func manhattanDistance(p1 Point, p2 Point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parseStringLineIntoArray(line string) []int64 {
	array := make([]int64, 0)
	for _, num := range strings.Split(strings.TrimSpace(line), " ") {
		n, _ := strconv.ParseInt(num, 10, 64)
		array = append(array, n)
	}
	return array
}
