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

	fileReader.Scan()
	line := fileReader.Text()
	parts := strings.Split(line, ":")
	times := parseStringLineIntoArray(strings.TrimSpace(parts[1]))

	fileReader.Scan()
	line = fileReader.Text()
	parts = strings.Split(line, ":")
	distances := parseStringLineIntoArray(strings.TrimSpace(parts[1]))

	fmt.Println("times", times)
	fmt.Println("distances", distances)

	mul := 1
	for i := 0; i < len(times); i++ {
		// fmt.Println("first time", findFirstTime(times[i], distances[i]))
		// fmt.Println("last time", findLastTime(times[i], distances[i]))

		total := findLastTimeLazy(times[i], distances[i]) - findFirstTimeLazy(times[i], distances[i]) + 1
		mul *= total
	}

	fmt.Println("mul", mul)

}

func parseStringLineIntoArray(line string) []int {
	array := make([]int, 0)
	for _, num := range strings.Split(strings.TrimSpace(line), " ") {
		if strings.TrimSpace(num) == "" {
			continue
		}
		n, _ := strconv.Atoi(num)
		array = append(array, n)
	}
	return array
}

func findFirstTimeBinarySearch(time int, distance int) int {
	left := 1
	right := time - 1
	for left <= right {
		mid := (left + right) / 2
		speed := mid
		if (time-mid)*speed > distance {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func findLastTimeBinarySearch(time int, distance int) int {
	left := 1
	right := time - 1
	for left <= right {
		mid := (left + right) / 2
		speed := mid
		if (time-mid)*speed > distance {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

func findFirstTimeLazy(time int, distance int) int {
	for i := 1; i < time; i++ {
		speed := i
		if (time-i)*speed > distance {
			return i
		}
	}
	return -1
}

func findLastTimeLazy(time int, distance int) int {
	for i := time - 1; i > 0; i-- {
		speed := i
		if (time-i)*speed > distance {
			return i
		}
	}
	return -1
}
