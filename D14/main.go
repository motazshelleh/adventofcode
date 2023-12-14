package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func main() {
	fd, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		panic(err)
	}
	defer fd.Close()

	fileReader := bufio.NewScanner(fd)

	arr := make([][]rune, 0)
	for fileReader.Scan() {
		line := fileReader.Text()
		arr = append(arr, []rune(line))
	}

	// copy arr
	part1Copy := make([][]rune, len(arr))
	for i := 0; i < len(arr); i++ {
		part1Copy[i] = make([]rune, len(arr[i]))
		copy(part1Copy[i], arr[i])
	}

	moveRocksNorthOrSouth(part1Copy, 1)
	fmt.Println("Part 1: ", calculateLoadONorth(part1Copy))

	m := make(map[string]int)
	iterations := 1000000000
	updated := false
	for i := 0; i < iterations; i++ {
		if lastDupIndex, ok := m[key(arr)]; !updated && ok {
			loopSize := i - lastDupIndex
			i += loopSize * ((iterations - loopSize - 1 - i) / loopSize)
			updated = true
		}
		m[key(arr)] = i
		moveRocksNorthOrSouth(arr, 1)
		moveRocksWestOrEast(arr, -1)
		moveRocksNorthOrSouth(arr, -1)
		moveRocksWestOrEast(arr, 1)
	}
	fmt.Println("Part 2: ", calculateLoadONorth(arr))

}

func key(arr [][]rune) string {
	temp := ""
	for _, a := range arr {
		temp += string(a)
	}
	return temp
}

func calculateLoadONorth(arr [][]rune) int {
	total := 0
	for i := 0; i < len(arr[0]); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[j][i] == 'O' {
				total += (len(arr) - j)
			}
		}
	}
	return total
}

func printArr(arr [][]rune) {
	for _, a := range arr {
		fmt.Println(string(a))
	}
}

// north = 1
// south = -1
func moveRocksNorthOrSouth(arr [][]rune, dir int) {
	for i := 0; i < len(arr[0]); i++ {
		emptySpacePtr, rockPtr := Point{len(arr) - 1, i}, Point{len(arr) - 1, i}
		if dir == 1 {
			emptySpacePtr, rockPtr = Point{0, i}, Point{0, i}
		}
		prevemptySpacePtr, prevRockPtr := Point{0, -1}, Point{0, -1}
		for isWithinBounds1(arr, rockPtr) && isWithinBounds1(arr, emptySpacePtr) {
			// fmt.Println("emptySpacePtr: ", emptySpacePtr, " rockPtr: ", rockPtr)
			if prevemptySpacePtr == emptySpacePtr && prevRockPtr == rockPtr {
				printArr(arr)
				panic("infinite loop")
			}
			prevemptySpacePtr = emptySpacePtr
			prevRockPtr = rockPtr

			for isWithinBounds1(arr, emptySpacePtr) && arr[emptySpacePtr.x][emptySpacePtr.y] != '.' {
				emptySpacePtr.x += dir
			}

			for isWithinBounds1(arr, rockPtr) && arr[rockPtr.x][rockPtr.y] == '.' {
				rockPtr.x += dir
			}

			if dir == 1 && emptySpacePtr.x > rockPtr.x || dir == -1 && emptySpacePtr.x < rockPtr.x {
				rockPtr.x = emptySpacePtr.x
				continue
			}

			if isWithinBounds1(arr, rockPtr) && isWithinBounds1(arr, emptySpacePtr) {
				if arr[rockPtr.x][rockPtr.y] == '#' {
					rockPtr.x += dir
					emptySpacePtr.x = rockPtr.x
					continue
				} else if arr[rockPtr.x][rockPtr.y] == 'O' {
					temp := arr[emptySpacePtr.x][emptySpacePtr.y]
					arr[emptySpacePtr.x][emptySpacePtr.y] = arr[rockPtr.x][rockPtr.y]
					arr[rockPtr.x][rockPtr.y] = temp
					if arr[rockPtr.x][rockPtr.y] != '.' {
						fmt.Println("invalid char1: ", arr[rockPtr.x][rockPtr.y])
						panic("invalid char1")
					}
					rockPtr.x += dir
				} else {
					panic("invalid char")
				}
			} else {
				break
			}
		}

	}
}

func moveRocksWestOrEast(arr [][]rune, dir int) {
	for i := 0; i < len(arr); i++ {
		emptySpacePtr, rockPtr := Point{i, 0}, Point{i, 0}
		if dir == 1 {
			emptySpacePtr, rockPtr = Point{i, len(arr) - 1}, Point{i, len(arr) - 1}
		}
		prevemptySpacePtr, prevRockPtr := Point{0, -1}, Point{0, -1}
		for isWithinBounds1(arr, rockPtr) && isWithinBounds1(arr, emptySpacePtr) {
			// fmt.Println("emptySpacePtr: ", emptySpacePtr, " rockPtr: ", rockPtr)
			if prevemptySpacePtr == emptySpacePtr && prevRockPtr == rockPtr {
				printArr(arr)
				panic("infinite loop")
			}
			prevemptySpacePtr = emptySpacePtr
			prevRockPtr = rockPtr

			for isWithinBounds1(arr, emptySpacePtr) && arr[emptySpacePtr.x][emptySpacePtr.y] != '.' {
				emptySpacePtr.y += -dir
			}

			for isWithinBounds1(arr, rockPtr) && arr[rockPtr.x][rockPtr.y] == '.' {
				rockPtr.y += -dir
			}

			if dir == 1 && emptySpacePtr.y < rockPtr.y || dir == -1 && emptySpacePtr.y > rockPtr.y {
				rockPtr.y = emptySpacePtr.y
				continue
			}

			if isWithinBounds1(arr, rockPtr) && isWithinBounds1(arr, emptySpacePtr) {
				if arr[rockPtr.x][rockPtr.y] == '#' {
					rockPtr.y += -dir
					emptySpacePtr.y = rockPtr.y
					continue
				} else if arr[rockPtr.x][rockPtr.y] == 'O' {
					temp := arr[emptySpacePtr.x][emptySpacePtr.y]
					arr[emptySpacePtr.x][emptySpacePtr.y] = arr[rockPtr.x][rockPtr.y]
					arr[rockPtr.x][rockPtr.y] = temp
					if arr[rockPtr.x][rockPtr.y] != '.' {
						fmt.Println("invalid char1: ", arr[rockPtr.x][rockPtr.y])
						panic("invalid char1")
					}
					rockPtr.y += -dir
				} else {
					panic("invalid char")
				}
			} else {
				break
			}
		}

	}
}

func isWithinBounds1(arr [][]rune, p Point) bool {
	return p.x < len(arr[0]) && p.y < len(arr) && p.x >= 0 && p.y >= 0
}
