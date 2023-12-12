package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var recursionCounter = 0

type Spring struct {
	field    []rune
	group    []int
	unknowns int
	groupSum int
}

func main() {
	fmt.Println("Solving part 1")
	solveMain(1)
	fmt.Println()
	fmt.Println("Solving part 2")
	solveMain(5)
}

func solveMain(expandTimes int) {
	recursionCounter = 0
	fd, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		panic(err)
	}
	defer fd.Close()

	fileReader := bufio.NewScanner(fd)

	springArr := make([]Spring, 0)
	for fileReader.Scan() {
		line := fileReader.Text()
		parts := strings.Split(line, " ")
		parts[0] = expandLine(parts[0], expandTimes, "?")
		parts[1] = expandLine(parts[1], expandTimes, ",")
		s := Spring{}
		s.field = []rune(parts[0])
		for _, c := range s.field {
			if c == '?' {
				s.unknowns++
			}
		}
		s.group, s.groupSum = parseStringLineIntoArray(parts[1])

		springArr = append(springArr, s)
	}

	total := 0
	for _, s := range springArr {
		sum := 0
		totalQuestionMarks := 0
		for _, n := range s.group {
			sum += n
		}
		for _, c := range s.field {
			if c == '#' {
				sum--
			}
			if c == '?' {
				totalQuestionMarks++
			}
		}
		s.unknowns = totalQuestionMarks
		// fmt.Println("processing: ", i, " Of: ", len(springArr))
		temp := solve(s, 0, s.unknowns, sum, 0, 0, make(map[string]int))

		total += temp
	}

	fmt.Println("Total: ", total)
	fmt.Println("Count: ", recursionCounter)
}

func expandLine(line string, times int, sep string) string {
	line1 := ""
	for i := 0; i < times; i++ {
		if i == 0 {
			line1 = line
		} else {
			line1 += sep + line
		}
	}
	return line1
}

func parseStringLineIntoArray(line string) ([]int, int) {
	array := make([]int, 0)
	sum := 0
	for _, num := range strings.Split(strings.TrimSpace(line), ",") {
		n, _ := strconv.Atoi(num)
		sum += n
		array = append(array, n)
	}
	return array, sum
}

func isSpringComplete(spring Spring) bool {
	currentSpringLength := 0
	groupIndex := 0
	for _, c := range spring.field {
		if c == '#' {
			currentSpringLength++
		} else if currentSpringLength > 0 {
			if groupIndex >= len(spring.group) || spring.group[groupIndex] != currentSpringLength {
				return false
			}
			groupIndex++
			currentSpringLength = 0
		}
	}
	if currentSpringLength > 0 {
		if groupIndex >= len(spring.group) || spring.group[groupIndex] != currentSpringLength {
			return false
		}
		groupIndex++
		currentSpringLength = 0
	}
	return groupIndex == len(spring.group) && currentSpringLength == 0
}

// walk through the field for each ? try to place a spring and recurse, and then try to place a dot and recurse
// keep track of the current group and the current group index
// if we try to create a new group (for example we reached a dot after spring) then we check if the number of the last group match the number of the expected group
// memorize the results of the recursion to avoid recalculating
func solve(spring Spring, fieldIndex int, remainingUnknowns int, remainingSprings int, groupIndex int, currentGroup int, visited map[string]int) int {
	key := strconv.Itoa(fieldIndex) + ":" + strconv.Itoa(groupIndex) + ":" + strconv.Itoa(currentGroup)
	if val, ok := visited[key]; ok {
		return val
	}
	recursionCounter++
	if remainingSprings == 0 {
		if isSpringComplete(spring) {
			return 1
		}
		return 0
	}
	if remainingUnknowns < remainingSprings {
		return 0
	}
	if fieldIndex >= len(spring.field) {
		return 0
	}
	totalSolutions := 0
	for i := fieldIndex; i < len(spring.field); i++ {
		if remainingSprings > len(spring.field)-i || groupIndex >= len(spring.group) || currentGroup > spring.group[groupIndex] {
			break
		}
		if spring.field[i] == '#' {
			currentGroup++
			continue
		}
		if spring.field[i] == '.' {
			if currentGroup != 0 && currentGroup != spring.group[groupIndex] {
				break
			}
			if currentGroup != 0 {
				currentGroup = 0
				groupIndex++
			}
			continue
		}
		if spring.field[i] == '?' {

			currentTotal := 0
			tempF := spring.field[i]
			spring.field[i] = '#'
			if currentGroup != 0 && spring.field[i] == '.' {
				if currentGroup != spring.group[groupIndex] {
					break
				}
				currentTotal += solve(spring, i+1, remainingUnknowns-1, remainingSprings, groupIndex+1, 0, visited)
			} else {
				currentTotal += solve(spring, i+1, remainingUnknowns-1, remainingSprings-1, groupIndex, currentGroup+1, visited)
			}

			spring.field[i] = '.'
			if currentGroup == 0 {
				currentTotal += solve(spring, i+1, remainingUnknowns-1, remainingSprings, groupIndex, 0, visited)
			} else if currentGroup == spring.group[groupIndex] {
				currentTotal += solve(spring, i+1, remainingUnknowns-1, remainingSprings, groupIndex+1, 0, visited)
			}

			spring.field[i] = tempF

			totalSolutions += currentTotal
		}
		break
	}
	visited[key] = totalSolutions
	return totalSolutions
}
