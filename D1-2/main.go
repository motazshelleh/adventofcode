package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// define a map with key as string and value as int, and fill it with numbers from 0 to 9
var numbersMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9"}

func main() {
	fd, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		panic(err)
	}
	defer fd.Close()
	sum := 0
	inputScanner := bufio.NewScanner(fd)
	for inputScanner.Scan() {
		line := inputScanner.Text()
		firstNum := findFirstNumber(line)
		lastNum := findLastNumber(line)
		compinedNums := firstNum + lastNum

		n, err := strconv.Atoi(compinedNums)
		if err != nil {
			fmt.Println("Error converting string to int " + compinedNums)
			panic(err)
		}
		sum += n
	}

	fmt.Println(sum)
}

func findFirstNumber(line string) string {
	value := ""
	numIndex := len(line)
	for i := 0; i < len(line); i++ {
		c := line[i]
		if c >= '0' && c <= '9' {
			numIndex = i
			value = string(c)
			break
		}
	}

	minIndex := len(line)
	for k, v := range numbersMap {
		if i := strings.Index(line[0:numIndex], k); i != -1 {
			if i < minIndex {
				minIndex = i
				value = v
			}
		}
	}

	return value
}

func findLastNumber(line string) string {
	value := ""
	numIndex := 0
	for i := len(line) - 1; i >= 0; i-- {
		c := line[i]
		if c >= '0' && c <= '9' {
			numIndex = i
			value = string(c)
			break
		}
	}

	maxIndex := 0
	for k, v := range numbersMap {
		if i := strings.LastIndex(line[numIndex:], k); i != -1 {
			if i > maxIndex {
				maxIndex = i
				value = v
			}
		}
	}

	return value
}
