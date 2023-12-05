package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type rangeMap struct {
	dest   int64
	src    int64
	length int64
}

func main() {
	fd, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		panic(err)
	}
	defer fd.Close()

	fileReader := bufio.NewScanner(fd)
	fileReader.Scan()
	seedsRawInput := fileReader.Text()[6:]
	var seeds []int64 = parseStringLineIntoArray(seedsRawInput)
	fmt.Println("seeds", seeds)
	rangeMaps := make([][]rangeMap, 0)
	currentRangeMap := make([]rangeMap, 0)
	for fileReader.Scan() {
		line := fileReader.Text()
		if strings.TrimSpace(line) == "" {
			if len(currentRangeMap) > 0 {
				sortRangeMap(currentRangeMap)
				rangeMaps = append(rangeMaps, currentRangeMap)
			}
			currentRangeMap = make([]rangeMap, 0)
			continue
		}
		if strings.HasSuffix(strings.TrimSpace(line), ":") {
			continue
		}
		rm := rangeMap{}
		parsedLine := parseStringLineIntoArray(line)
		rm.dest = parsedLine[0]
		rm.src = parsedLine[1]
		rm.length = parsedLine[2]
		currentRangeMap = append(currentRangeMap, rm)
	}
	rangeMaps = append(rangeMaps, currentRangeMap)

	src := seeds
	for _, rm := range rangeMaps {
		sort.Slice(src, func(i, j int) bool {
			return src[i] < src[j]
		})
		src = mapFromSourceToDest(rm, src)
	}
	fmt.Println("locations:", src)
	min := src[0]
	for _, v := range src {
		if v < min {
			min = v
		}
	}
	fmt.Println("min", min)
}

func parseStringLineIntoArray(line string) []int64 {
	array := make([]int64, 0)
	for _, num := range strings.Split(strings.TrimSpace(line), " ") {
		n, _ := strconv.ParseInt(num, 10, 64)
		array = append(array, n)
	}
	return array
}

func sortRangeMap(rm []rangeMap) {
	sort.Slice(rm, func(i, j int) bool {
		return rm[i].src < rm[j].src
	})
}

func mapFromSourceToDest(rm []rangeMap, input []int64) []int64 {
	output := make([]int64, len(input))
	for i := 0; i < len(input); i++ {
		ok, r := validate(rm, input[i])
		if ok {
			mappedValue := input[i] - r.src + r.dest
			output[i] = mappedValue
		} else {
			output[i] = input[i]
		}
	}

	return output
}

func validate(rm []rangeMap, input int64) (bool, rangeMap) {
	for _, v := range rm {
		if input >= v.src && input <= v.src+v.length-1 {
			return true, v
		}
	}
	return false, rangeMap{}
}
