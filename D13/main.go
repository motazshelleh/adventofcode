package main

import (
	"bufio"
	"fmt"
	"os"
)

type Field struct {
	horizontal    []string
	vertical      []string
	horizontalMap map[string][]int
	verticalMap   map[string][]int
}

func main() {
	fd, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		panic(err)
	}
	defer fd.Close()

	fileReader := bufio.NewScanner(fd)

	fields := make([]Field, 0)
	f := Field{}
	f.horizontal = make([]string, 0)
	for fileReader.Scan() {
		line := fileReader.Text()
		if line == "" {
			fields = append(fields, f)
			f = Field{}
			f.horizontal = make([]string, 0)
		} else {
			f.horizontal = append(f.horizontal, line)
		}
	}
	fields = append(fields, f)

	// for _, f := range fields {
	// 	for _, h := range f.horizontal {
	// 		fmt.Println(h)
	// 	}
	// }

	for ind, f := range fields {
		f.verticalMap = make(map[string][]int)
		f.horizontalMap = make(map[string][]int)
		f.vertical = make([]string, 0)
		for i, h := range f.horizontal {
			f.horizontalMap[h] = append(f.horizontalMap[h], i)
		}

		for i := 0; i < len(f.horizontal[0]); i++ {
			temp := ""
			for _, h := range f.horizontal {
				temp += string(h[i])
			}
			f.vertical = append(f.vertical, temp)
			f.verticalMap[temp] = append(f.verticalMap[temp], i)
		}
		fields[ind] = f
	}

	part1 := 0
	for _, f := range fields {
		h := findMirror(f.horizontal)
		v := findMirror(f.vertical)

		if h > v {
			part1 += (h * 100)
		} else {
			part1 += v
		}
	}
	fmt.Println("part1 : ", part1)

	part2 := 0
	for _, f := range fields {
		h := solve(f.horizontal, f.horizontalMap)
		v := solve(f.vertical, f.verticalMap)

		if h > v {
			part2 += (h * 100)
		} else {
			part2 += v
		}
	}
	fmt.Println("part2 : ", part2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(line []string, m map[string][]int) int {
	beforeFixIndex, beforeFixSize := findMirrorWithDefect(line, m, -1)
	for i, l := range line {
		copy := []rune(l)
		for ci := 0; ci < len(l); ci++ {
			copy = flip(copy, ci)
			boo1, boo2 := false, false
			if _, ok := m[string(copy)]; ok {
				m[string(copy)] = addValueToSortedSlice(m[string(copy)], i)
				boo1 = true
			}
			if arr, ok := m[l]; ok {
				if len(arr) > 1 {
					m[l] = removeValueFromSlice(m[l], i)
					boo2 = true
				}
			}
			if boo1 || boo2 {
				temp := l
				line[i] = string(copy)
				val, valSize := findMirrorWithDefect(line, m, beforeFixIndex)
				line[i] = temp
				if val != -1 && (val != beforeFixIndex || beforeFixSize != valSize) {
					if boo1 {
						m[string(copy)] = removeValueFromSlice(m[string(copy)], i)
					}
					if boo2 {
						m[l] = addValueToSortedSlice(m[l], i)
					}

					return val
				}
			}
			if boo1 {
				m[string(copy)] = removeValueFromSlice(m[string(copy)], i)
			}
			if boo2 {
				m[l] = addValueToSortedSlice(m[l], i)
			}
			copy = flip(copy, ci)
		}
	}
	return -1
}

func removeValueFromSlice(s []int, v int) []int {
	for i, val := range s {
		if val == v {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func addValueToSortedSlice(s []int, v int) []int {
	for i, val := range s {
		if val > v {
			return append(s[:i], append([]int{v}, s[i:]...)...)
		}
	}
	return append(s, v)
}

func flip(r []rune, i int) []rune {
	if r[i] == '#' {
		r[i] = '.'
	} else {
		r[i] = '#'
	}
	return r
}

func printDebug(line []string, m map[string][]int) {
	for _, l := range line {
		fmt.Println(l, m[l])
	}
	fmt.Println()
}

func findMirrorWithDefect(line []string, m map[string][]int, ignoreIndex int) (int, int) {
	// printDebug(line)
	maxSize := 0
	maxIndex := -1
	for _, v := range m {
		for i := 0; i < len(v)-1; i++ {
			if v[i]+1 == v[i+1] {
				size, edge := MaxMirrorSize(line, v[i+1])
				if edge {
					if ignoreIndex != v[i+1] || ignoreIndex == -1 {
						maxSize = size
						maxIndex = v[i+1]
					}
					// return v[i+1]
				}
			}
		}
	}
	return maxIndex, maxSize
}

func findMirror(line []string) int {
	indexOfMirror := -1
	for i := 0; i < len(line)-1; i++ {
		if line[i] == line[i+1] {
			_, edge := MaxMirrorSize(line, i+1)
			if edge {
				indexOfMirror = i + 1
				break
			}
		}
	}
	return indexOfMirror
}

func isDiffOnlyOne(line1, line2 string) bool {
	diffCount := 0
	for i := 0; i < len(line1); i++ {
		if line1[i] != line2[i] {
			diffCount++
			if diffCount > 1 {
				return false
			}
		}
	}
	return diffCount == 1
}

func MaxMirrorSize(line []string, start int) (int, bool) {
	p1 := start
	p2 := start - 1
	count := 0
	for p1 >= 0 && p2 < len(line) {
		if line[p1] == line[p2] {
			count++
		} else {
			break
		}
		p1--
		p2++
	}
	if p1 < 0 || p2 >= len(line) {
		return count, true
	}
	return count, false
}
