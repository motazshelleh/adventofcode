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

	finalMin := int64(-1)
	for i := 0; i < len(seeds); i += 2 {
		fmt.Println("...")
		src := make([]int64, 0)
		for j := int64(seeds[i]); j < seeds[i]+seeds[i+1]; j++ {
			src = append(src, j)
			// batch size is 1 million
			if len(src) < 1000000 && j < seeds[i]+seeds[i+1]-1 {
				continue
			}
			min := int64(999999999999999999)
			// fmt.Println("src", src)
			for _, rm := range rangeMaps {
				// for _, v := range rm {
				// 	fmt.Println(v)
				// }
				// sort.Slice(src, func(i, j int) bool {
				// 	return src[i] < src[j]
				// })
				min, src = mapFromSourceToDest(rm, src)
				// fmt.Println("src", src)
				// break
			}

			for _, v := range src {
				if v < min {
					min = v
				}
			}
			if finalMin == -1 || min < finalMin {
				finalMin = min
			}

			src = make([]int64, 0)
		}
	}

	fmt.Println("min", finalMin)
}

func parseStringLineIntoArray(line string) []int64 {
	array := make([]int64, 0)
	for _, num := range strings.Split(strings.TrimSpace(line), " ") {
		n, _ := strconv.ParseInt(num, 10, 64)
		array = append(array, n)
	}
	return array
}

func expandSeedsArray(seeds []int64) []int64 {
	expanded := make([]int64, 0)
	fmt.Println("seeds", seeds)
	for i := 0; i < len(seeds); i += 2 {
		for j := int64(seeds[i]); j < seeds[i]+seeds[i+1]; j++ {
			expanded = append(expanded, j)
		}
	}
	return expanded
}

func sortRangeMap(rm []rangeMap) {
	sort.Slice(rm, func(i, j int) bool {
		return rm[i].src < rm[j].src
	})
}

func mapFromSourceToDest(rm []rangeMap, input []int64) (int64, []int64) {
	output := make([]int64, len(input))
	// fmt.Println("rm", rm)
	// rangeMapIndex := 0
	minValue := int64(-1)
	for i := 0; i < len(input); i++ {
		// fmt.Println(rangeMapIndex)

		// for rangeMapIndex < len(rm) && input[i] > rm[rangeMapIndex].src+rm[rangeMapIndex].length-1 {
		// 	rangeMapIndex++
		// }

		// fmt.Println("rangeMapIndex", rangeMapIndex, "input[i]", input[i])
		// if rangeMapIndex < len(rm) && input[i] >= rm[rangeMapIndex].src && input[i] <= rm[rangeMapIndex].src+rm[rangeMapIndex].length-1 {
		ok, r := validate(rm, input[i])
		if ok {
			mappedValue := input[i] - r.src + r.dest
			// fmt.Println("mappedValue", mappedValue)
			if minValue == -1 || mappedValue < minValue {
				minValue = mappedValue
			}

			output[i] = mappedValue
		} else {

			if minValue == -1 || input[i] < minValue {
				minValue = input[i]
			}

			output[i] = input[i]
		}
	}

	return minValue, output
}

func validate(rm []rangeMap, input int64) (bool, rangeMap) {
	for _, v := range rm {
		if input >= v.src && input <= v.src+v.length-1 {
			return true, v
		}
	}
	return false, rangeMap{}
}
