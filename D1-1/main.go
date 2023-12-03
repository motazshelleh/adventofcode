package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
		firstNum := findFirstNumberIndex(line)
		lastNum := findLastNumberIndex(line)
		compinedNums := string(line[firstNum]) + string(line[lastNum])

		n, err := strconv.Atoi(compinedNums)
		if err != nil {
			fmt.Println("Error converting string to int " + compinedNums)
			panic(err)
		}
		sum += n
	}

	fmt.Println(sum)
}

func findFirstNumberIndex(line string) int {
	for i := 0; i < len(line); i++ {
		c := line[i]
		if c >= '0' && c <= '9' {
			return i
		}
	}
	return -1
}

func findLastNumberIndex(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		c := line[i]
		if c >= '0' && c <= '9' {
			return i
		}
	}
	return -1
}
