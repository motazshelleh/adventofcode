package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fd, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		panic(err)
	}

	arr := make([][]rune, 0)

	fileReader := bufio.NewScanner(fd)
	for fileReader.Scan() {
		line := fileReader.Text()
		tempArr := make([]rune, 0)
		for i := 0; i < len(line); i++ {
			tempArr = append(tempArr, rune(line[i]))
		}
		arr = append(arr, tempArr)
	}

	sum := 0
	for i := 0; i < len(arr); i++ {
		numStr := ""
		isNextToSymbol := false
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] >= '0' && arr[i][j] <= '9' {
				numStr += string(arr[i][j])
				if isAdjacentToSymbol(arr, i, j) {
					isNextToSymbol = true
				}
			} else {
				sum += enginePartNumber(numStr, isNextToSymbol)
				numStr = ""
				isNextToSymbol = false
			}
		}
		sum += enginePartNumber(numStr, isNextToSymbol)
	}

	fmt.Println(sum)
}

func enginePartNumber(numStr string, isNextToSymbol bool) int {
	num := 0
	if len(numStr) > 0 && isNextToSymbol {
		fmt.Sscanf(numStr, "%d", &num)
	}
	return num
}

func isAdjacentToSymbol(arr [][]rune, i int, j int) bool {
	deltaOffsets := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	for _, deltaOffset := range deltaOffsets {
		if isInBounds(arr, i+deltaOffset[0], j+deltaOffset[1]) && isSymbol(arr[i+deltaOffset[0]][j+deltaOffset[1]]) {
			return true
		}
	}
	return false
}

func isInBounds(arr [][]rune, i int, j int) bool {
	if i >= 0 && i < len(arr) && j >= 0 && j < len(arr[i]) {
		return true
	}
	return false
}

func isSymbol(c rune) bool {
	if (c < '0' || c > '9') && c != '.' {
		return true
	}
	return false
}
