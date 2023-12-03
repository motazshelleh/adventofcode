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
	defer fd.Close()
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

	gearMap := make(map[string][]int)
	for i := 0; i < len(arr); i++ {
		numStr := ""
		isPrevNumber := false
		gearPoss := make([][]int, 0)
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] >= '0' && arr[i][j] <= '9' {
				numStr += string(arr[i][j])
				gearPos := isAdjacentToSymbol(arr, i, j, !isPrevNumber, !isNumber(arr, i, j+1))
				gearPoss = append(gearPoss, gearPos...)
				isPrevNumber = true
			} else {
				addGearMap(gearMap, gearPoss, numStr)
				isPrevNumber = false
				numStr = ""
				gearPoss = make([][]int, 0)
			}
		}
		addGearMap(gearMap, gearPoss, numStr)
	}

	sum := 0
	for _, gearNums := range gearMap {
		if len(gearNums) == 2 {
			mul := gearNums[0] * gearNums[1]
			sum += mul
		}
	}
	fmt.Println(sum)
}

func addGearMap(gearMap map[string][]int, gearPoss [][]int, numStr string) {
	if len(numStr) > 0 {
		num := 0
		fmt.Sscanf(numStr, "%d", &num)
		for _, pos := range gearPoss {
			posIndex := fmt.Sprintf("%d->%d", pos[0], pos[1])
			gearMap[posIndex] = append(gearMap[posIndex], num)
		}
	}
}

func isAdjacentToSymbol(arr [][]rune, i int, j int, isFirstNumber bool, isLastNumber bool) [][]int {
	ansArr := make([][]int, 0)
	deltaOffsets := make([][]int, 0)
	deltaOffsets = append(deltaOffsets, []int{-1, 0}, []int{1, 0})
	if isFirstNumber {
		deltaOffsets = append(deltaOffsets, []int{-1, -1}, []int{0, -1}, []int{1, -1})
	}
	if isLastNumber {
		deltaOffsets = append(deltaOffsets, []int{-1, 1}, []int{0, 1}, []int{1, 1})
	}

	for _, deltaOffset := range deltaOffsets {
		if isInBounds(arr, i+deltaOffset[0], j+deltaOffset[1]) && isSymbol(arr[i+deltaOffset[0]][j+deltaOffset[1]]) {
			ansArr = append(ansArr, []int{i + deltaOffset[0], j + deltaOffset[1]})
		}
	}
	return ansArr
}

func isInBounds(arr [][]rune, i int, j int) bool {
	if i >= 0 && i < len(arr) && j >= 0 && j < len(arr[i]) {
		return true
	}
	return false
}

func isSymbol(c rune) bool {
	if c == '*' {
		return true
	}
	return false
}

func isNumber(arr [][]rune, i int, j int) bool {
	if isInBounds(arr, i, j) && arr[i][j] >= '0' && arr[i][j] <= '9' {
		return true
	}
	return false
}
