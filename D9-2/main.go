package main

import (
	"bufio"
	"fmt"
	"os"
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

	sum := 0
	for fileReader.Scan() {
		line := fileReader.Text()
		history := parseStringLineIntoArray(line)
		missingValue := findTheMissingValue(history)
		fmt.Println("missingValue", missingValue)
		sum += missingValue
		// fmt.Print(".")
	}
	fmt.Println(".")

	fmt.Println("sum", sum)

}

func parseStringLineIntoArray(line string) []int {
	array := make([]int, 0)
	for _, num := range strings.Split(strings.TrimSpace(line), " ") {
		n, _ := strconv.Atoi(num)
		array = append(array, n)
	}
	return array
}

func findTheMissingValue(history []int) int {
	arr := make([][]int, 0)
	allZeros := false
	arr = append(arr, history)
	for !allZeros {
		allZeros = true
		tempArr := make([]int, len(history)-1)
		for i := 0; i < len(history)-1; i++ {
			tempArr[i] = history[i+1] - history[i]
			if tempArr[i] != 0 {
				allZeros = false
			}
		}
		if len(tempArr) == 0 {
			arr = append(arr, []int{0})
		} else {
			arr = append(arr, tempArr)
		}

		history = tempArr
	}

	// fmt.Println("arr", arr)

	prevNumber := 0
	for i := len(arr) - 1; i >= 0; i-- {
		// fmt.Println(arr[i][len(arr[i])-1], prevNumber)
		prevNumber = arr[i][len(arr[i])-1] + prevNumber
		// fmt.Println("prevNumber", prevNumber)
	}
	return prevNumber
}
