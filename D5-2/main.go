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
	from     int64
	to       int64
	mappedTo int64
}

type seed struct {
	from int64
	to   int64
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

	seedsArr := make([]seed, 0)

	for i := 0; i < len(seeds); i += 2 {
		seedsArr = append(seedsArr, seed{seeds[i], seeds[i] + seeds[i+1] - 1})
	}

	fmt.Println("seeds", seedsArr)
	rangeMaps := make([][]rangeMap, 0)
	currentRangeMap := make([]rangeMap, 0)
	for fileReader.Scan() {
		line := fileReader.Text()
		if strings.TrimSpace(line) == "" {
			if len(currentRangeMap) > 0 {
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
		rm.from = parsedLine[1]
		rm.to = parsedLine[1] + parsedLine[2] - 1
		rm.mappedTo = parsedLine[0]
		currentRangeMap = append(currentRangeMap, rm)
	}
	rangeMaps = append(rangeMaps, currentRangeMap)

	for i := 0; i < len(rangeMaps); i++ {
		// first we fill the missing rangeMaps
		// second we sort -> expand & map the seeds
		// third after going through all rangeMap we take the min value from the seeds
		rangeMaps[i] = fillMissingRangeMap(rangeMaps[i])
		seedsArr = expandRangeMap(rangeMaps[i], seedsArr)
	}

	min := int64(-1)
	for _, v := range seedsArr {
		if min == -1 || v.from < min {
			min = v.from
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
		return rm[i].from < rm[j].from
	})
}

func fillMissingRangeMap(rm []rangeMap) []rangeMap {
	sortRangeMap(rm)
	filled := make([]rangeMap, 0)
	if rm[0].from > 0 {
		filled = append(filled, rangeMap{0, rm[0].from - 1, 0})
	}
	for i := 0; i < len(rm)-1; i++ {
		filled = append(filled, rm[i])
		if rm[i].to < rm[i+1].from-1 {
			filled = append(filled, rangeMap{rm[i].to + 1, rm[i+1].from - 1, rm[i].to + 1})
		}
	}
	filled = append(filled, rm[len(rm)-1])
	return filled
}

func expandRangeMap(rm []rangeMap, seeds []seed) []seed {
	expanded := make([]seed, 0)

	for i := 0; i < len(seeds); i++ {
		// for each seed we go through the rangeMap and check the boarders and map them to their correct location
		doneProcessingSeed := false
		for _, v := range rm {
			// until we finish rm or seeds[i] is fully expanded
			temp := seed{}
			// range.from is "source range start"
			// range.to is "source range start" + "range length"
			// check if the seed is within the rangeMap =>   |range.from    seed.from    range.to|
			if v.from <= seeds[i].from && v.to >= seeds[i].from {
				// if the seed is within the rangeMap
				// then we create a new seed with the mappedTo value
				diff := (seeds[i].from - v.from)
				temp.from = v.mappedTo + diff

				// check if the seed is fully within the rangeMap =>   |range.from    seed.from    seed.to    range.to|
				if seeds[i].to <= v.to {
					// if the seed is fully within the rangeMap
					// then we close the seed at the relative location to mappedTo with the same original length and mark it as done
					temp.to = temp.from + (seeds[i].to - seeds[i].from)
					doneProcessingSeed = true
				} else {
					// if the seed is not fully within the rangeMap
					// then we close the seed at the end of the current range and the length is the remaining length till the range.to
					temp.to = temp.from + (v.to - seeds[i].from)
					seeds[i].from = v.to + 1
				}
				expanded = append(expanded, temp)
			}
			if doneProcessingSeed {
				break
			}
		}

		// add the remaining seeds if there is any
		if !doneProcessingSeed {
			expanded = append(expanded, seeds[i])
		}
	}
	return expanded
}
