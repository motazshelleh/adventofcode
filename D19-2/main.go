package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Condition struct {
	accept    bool
	reject    bool
	compareTo rune
	op        rune
	value     int
	dist      string
}

type Range struct {
	min int
	max int
}

func main() {
	fd, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		panic(err)
	}
	defer fd.Close()

	fileReader := bufio.NewScanner(fd)

	m := make(map[string][]Condition)
	for fileReader.Scan() {
		line := fileReader.Text()
		if len(line) == 0 {
			break
		}
		conditions, name := parseCondition(line)
		m[name] = conditions
	}

	input := make(map[rune]Range)
	input['x'] = Range{1, 4000}
	input['m'] = Range{1, 4000}
	input['a'] = Range{1, 4000}
	input['s'] = Range{1, 4000}

	total := solve(m, input, "in")
	fmt.Println(total)
}

func solve(m map[string][]Condition, input map[rune]Range, currentConditionKey string) int {
	currentCondition := m[currentConditionKey]
	localTotal := 0
	for _, condition := range currentCondition {
		if condition.compareTo != 0 {
			if condition.op == '>' {
				copy := copyRange(input)
				temp := copy[condition.compareTo]
				temp.max = min(temp.max, condition.value)
				copy[condition.compareTo] = temp

				temp = input[condition.compareTo]
				temp.min = max(temp.min, condition.value+1)
				input[condition.compareTo] = temp

				if input[condition.compareTo].min <= input[condition.compareTo].max && input[condition.compareTo].min > condition.value {
					if condition.accept {
						localTotal += getTotalCombination(input)
					}
					localTotal += solve(m, input, condition.dist)
				}
				input = copy
				if input[condition.compareTo].min > input[condition.compareTo].max {
					break
				}
			} else if condition.op == '<' {
				copy := copyRange(input)
				temp := copy[condition.compareTo]
				temp.min = max(temp.min, condition.value)
				copy[condition.compareTo] = temp

				temp = input[condition.compareTo]
				temp.max = min(temp.max, condition.value-1)
				input[condition.compareTo] = temp

				if input[condition.compareTo].min <= input[condition.compareTo].max && input[condition.compareTo].max < condition.value {
					if condition.accept {
						localTotal += getTotalCombination(input)
					}
					localTotal += solve(m, input, condition.dist)
				}
				input = copy
				if input[condition.compareTo].min > input[condition.compareTo].max {
					break
				}
			}
		} else {
			if condition.accept {
				localTotal += getTotalCombination(input)
			} else {
				localTotal += solve(m, input, condition.dist)
			}
			break
		}
	}
	return localTotal
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func copyRange(input map[rune]Range) map[rune]Range {
	output := make(map[rune]Range)
	for key, val := range input {
		r := Range{val.min, val.max}
		output[key] = r
	}
	return output
}

func getTotalCombination(input map[rune]Range) int {
	total := 1
	for _, val := range input {
		total *= val.max - val.min + 1
	}
	return total
}

func parseCondition(line string) ([]Condition, string) {
	parts := strings.Split(line, "{")
	name := parts[0]
	conditions := make([]Condition, 0)
	for _, part := range strings.Split(parts[1][0:len(parts[1])-1], ",") {
		c := Condition{}
		sides := strings.Split(part, ":")
		if len(sides) == 1 {
			if sides[0] == "A" {
				c.accept = true
			} else if sides[0] == "R" {
				c.reject = true
			} else {
				c.dist = sides[0]
			}
		} else {
			c.compareTo = rune(sides[0][0])
			c.op = rune(sides[0][1])
			c.value, _ = strconv.Atoi(sides[0][2:])
			if sides[1] == "A" {
				c.accept = true
			} else if sides[1] == "R" {
				c.reject = true
			} else {
				c.dist = sides[1]
			}
		}
		conditions = append(conditions, c)
	}
	return conditions, name
}
