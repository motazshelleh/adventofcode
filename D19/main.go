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

	inputList := make([]map[rune]int, 0)
	for fileReader.Scan() {
		line := fileReader.Text()
		line = line[1 : len(line)-1]
		input := make(map[rune]int)
		for _, part := range strings.Split(line, ",") {
			parts := strings.Split(part, "=")
			input[rune(parts[0][0])], _ = strconv.Atoi(parts[1])
		}
		inputList = append(inputList, input)
	}

	// fmt.Println(inputList)
	total := 0
	for _, input := range inputList {
		if solve(m, input, "in") {
			for _, val := range input {
				total += val
			}
		}
	}
	fmt.Println("Total:", total)
}

func solve(m map[string][]Condition, input map[rune]int, currentConditionKey string) bool {
	// fmt.Println(currentConditionKey, m[currentConditionKey])
	currentCondition := m[currentConditionKey]
	for _, condition := range currentCondition {
		if condition.compareTo != 0 {
			if condition.op == '>' {
				if input[condition.compareTo] > condition.value {
					if condition.accept {
						return true
					} else if condition.reject {
						return false
					}
					return solve(m, input, condition.dist)
				}
			} else if condition.op == '<' {
				if input[condition.compareTo] < condition.value {
					if condition.accept {
						return true
					} else if condition.reject {
						return false
					}
					return solve(m, input, condition.dist)
				}
			}
		} else {
			if condition.accept {
				return true
			} else if condition.reject {
				return false
			} else {
				return solve(m, input, condition.dist)
			}
		}
	}
	fmt.Println("No solution found", currentConditionKey)
	panic("No solution found")
	// return true
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
