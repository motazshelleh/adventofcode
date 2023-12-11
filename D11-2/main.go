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

var GALAXY_EXPANSION = 1000000

func main() {
	fd, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		panic(err)
	}
	defer fd.Close()

	// we remove 1 because we don't remove the original line so if GALAXY_EXPANSION is 10 then we need to add 9 lines
	GALAXY_EXPANSION -= 1

	fileReader := bufio.NewScanner(fd)

	inputMap := make([][]rune, 0)
	xExpansion := make([]int, 0)
	xExpansionLastValue := 0
	for fileReader.Scan() {
		line := fileReader.Text()
		inputMap = append(inputMap, []rune(line))
		isAllDots := true
		for i := 0; i < len(line); i++ {
			if line[i] != '.' {
				isAllDots = false
				break
			}
		}
		if isAllDots {
			xExpansionLastValue += GALAXY_EXPANSION
			xExpansion = append(xExpansion, xExpansionLastValue)
		} else {
			xExpansion = append(xExpansion, xExpansionLastValue)
		}
	}
	// for i := 0; i < len(inputMap); i++ {
	// 	fmt.Println(string(inputMap[i]))
	// }

	yExpansion := make([]int, 0)
	yExpansionLastValue := 0
	for i := 0; i < len(inputMap[0]); i++ {
		isAllDots := true
		for j := 0; j < len(inputMap); j++ {
			if inputMap[j][i] != '.' {
				isAllDots = false
			}
		}
		if isAllDots {
			yExpansionLastValue += GALAXY_EXPANSION
			yExpansion = append(yExpansion, yExpansionLastValue)
		} else {
			yExpansion = append(yExpansion, yExpansionLastValue)
		}
	}

	// }

	// for i := 0; i < len(inputMapExpanded); i++ {
	// 	fmt.Println(string(inputMapExpanded[i]))
	// }

	glaxiesPosition := make([]Point, 0)
	fmt.Println(yExpansion, xExpansion)
	for i := 0; i < len(inputMap); i++ {
		for j := 0; j < len(inputMap[i]); j++ {
			if inputMap[i][j] != '.' {
				p := Point{j + yExpansion[j], i + xExpansion[i]}
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
