package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fd, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		panic(err)
	}
	defer fd.Close()

	fileReader := bufio.NewScanner(fd)
	totalPoints := 0
	pointsArr := make([]int, 0)
	for fileReader.Scan() {
		line := fileReader.Text()
		chosenNums, winningNums := splitLine(line)
		points := findWinningPoints(chosenNums, winningNums)
		cardDuplicate := 0
		end := len(pointsArr)
		for i := 0; i < end; i++ {
			if pointsArr[i] != 0 {
				cardDuplicate += 1
				pointsArr[i]--
				pointsArr = append(pointsArr, points)
			}
		}
		pointsArr = append(pointsArr, points)
		totalPoints += cardDuplicate + 1
	}
	fmt.Println(totalPoints)
}

func splitLine(line string) ([]int, []int) {
	cardAndNumsSplit := strings.Split(line, ":")
	numsSplit := strings.Split(cardAndNumsSplit[1], "|")
	chosenNums := make([]int, 0)

	re := regexp.MustCompile("\\s+")
	for _, numStr := range re.Split(strings.TrimSpace(numsSplit[0]), -1) {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		chosenNums = append(chosenNums, num)
	}

	winningNums := make([]int, 0)
	for _, numStr := range re.Split(strings.TrimSpace(numsSplit[1]), -1) {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		winningNums = append(winningNums, num)
	}
	return chosenNums, winningNums
}

func findWinningPoints(chosenNums []int, winningNums []int) int {
	numbersMap := make(map[int]int)
	for _, num := range chosenNums {
		numbersMap[num] += 1
	}
	points := 0
	for _, num := range winningNums {
		if numbersMap[num] > 0 {
			points += 1
		}
	}
	return points
}
