package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lens struct {
	focalLength int
	value       string
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
	line := fileReader.Text()
	arr := strings.Split(line, ",")
	// part1(arr)

	m := make(map[int][]Lens)
	for _, a := range arr {
		addOp := true
		parts := strings.Split(a, "=")
		l := Lens{}
		if len(parts) == 1 {
			addOp = false
			parts = strings.Split(a, "-")
			l.value = parts[0]
		} else {
			l.focalLength, _ = strconv.Atoi(parts[1])
			l.value = parts[0]
		}
		hash := findHash(parts[0])
		// fmt.Println("hash:", hash, l)
		if addOp {
			addOrReplace(m, hash, l)
		} else {
			removeLens(m, hash, l)
		}
	}

	// debug(m)

	focusingPower := 0
	for boxNumber, v := range m {
		for lenIndex, l := range v {
			focusingPower += (1 + boxNumber) * (lenIndex + 1) * l.focalLength
		}
	}
	fmt.Println("Part 2: ", focusingPower)

}

func debug(m map[int][]Lens) {
	fmt.Println()
	for k, v := range m {
		if len(v) > 0 {
			fmt.Println(k, v)
		}
	}
	fmt.Println()
}

func removeLens(m map[int][]Lens, key int, lens Lens) {
	if _, ok := m[key]; ok {
		for i, l := range m[key] {
			if l.value == lens.value {
				m[key] = append(m[key][:i], m[key][i+1:]...)
				return
			}
		}
	}
}

func addOrReplace(m map[int][]Lens, key int, lens Lens) {
	if _, ok := m[key]; !ok {
		m[key] = append(m[key], lens)
	}
	for i, l := range m[key] {
		if l.value == lens.value {
			m[key][i] = lens
			return
		}
	}
	m[key] = append(m[key], lens)
}

func part1(arr []string) {
	total := 0
	for _, a := range arr {
		total += findHash(a)
	}
	fmt.Println("Part 1: ", total)
}

func findHash(str string) int {
	hash := 0
	for _, c := range str {
		hash += int(c)
		hash *= 17
		hash %= 256
	}
	return hash
}
