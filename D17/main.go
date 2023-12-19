package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Point struct {
	x            int
	y            int
	dir          rune
	visited      *map[string]int
	straightLine int
	pathValue    int
}

func main() {
	clearFile()
	fd, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		panic(err)
	}
	defer fd.Close()

	fileReader := bufio.NewScanner(fd)

	m := make([][]int, 0)
	for fileReader.Scan() {
		line := fileReader.Text()
		larr := make([]int, 0)
		for _, r := range line {
			larr = append(larr, int(r-'0'))
		}
		m = append(m, larr)
	}

	// debug(m)

	// p := Point{0, 0, 'R'}
	// minValue := solve(m, p, 0, make(map[string]int), make(map[string]bool))

	// fmt.Println("part1: ", minValue)
	fmt.Println("part1: ", solve(m))
}

func debug(m [][]int) {
	for _, row := range m {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

var count int = 0

func solve(m [][]int) int {
	// queue := make([]Point, 0)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	// queue = append(queue, Point{0, 0, 'R', make(map[string]bool), 0, 0})
	globalVis := make(map[string]int)
	tempp := &Point{0, 0, 'R', nil, 0, 0}
	heap.Push(&pq, tempp)
	minValue := 10000000000000000
	count := 0
	// for len(queue) > 0 {
	for pq.Len() > 0 {
		count++
		// sort.Slice(queue, func(i, j int) bool {
		// 	return queue[i].pathValue < queue[j].pathValue
		// })
		// point := queue[0]
		point := heap.Pop(&pq).(*Point)
		// debugQueue(queue)
		// queue = queue[1:]
		// if count == 10 {
		// 	panic("count exceeded")
		// }
		// fmt.Println(point.pathValue, pq.Len(), len(globalVis))

		if point.x == len(m)-1 && point.y == len(m[0])-1 {
			if point.straightLine < minValue {
				minValue = point.pathValue
				// panic("minValue: " + fmt.Sprintf("%d", minValue))
				return minValue
			}
			continue
		}

		key := getKey(*point, point.straightLine)
		if val, ok := globalVis[key]; ok && point.pathValue < val {
			globalVis[key] = point.pathValue
		} else if !ok {
			globalVis[key] = point.pathValue
		} else if ok {
			continue
		}
		// (*point.visited)[key] = point.pathValue

		deltas := make([]Point, 0)
		if point.dir == 'L' {
			deltas = append(deltas, Point{0, -1, 'L', nil, point.straightLine + 1, 0})
			deltas = append(deltas, Point{1, 0, 'D', nil, 0, 0})
			deltas = append(deltas, Point{-1, 0, 'U', nil, 0, 0})
		} else if point.dir == 'R' {
			deltas = append(deltas, Point{0, 1, 'R', nil, point.straightLine + 1, 0})
			deltas = append(deltas, Point{1, 0, 'D', nil, 0, 0})
			deltas = append(deltas, Point{-1, 0, 'U', nil, 0, 0})
		} else if point.dir == 'U' {
			deltas = append(deltas, Point{0, -1, 'L', nil, 0, 0})
			deltas = append(deltas, Point{0, 1, 'R', nil, 0, 0})
			deltas = append(deltas, Point{-1, 0, 'U', nil, point.straightLine + 1, 0})
		} else if point.dir == 'D' {
			deltas = append(deltas, Point{0, -1, 'L', nil, 0, 0})
			deltas = append(deltas, Point{0, 1, 'R', nil, 0, 0})
			deltas = append(deltas, Point{1, 0, 'D', nil, point.straightLine + 1, 0})
		}

		for _, delta := range deltas {
			p := Point{point.x + delta.x, point.y + delta.y, delta.dir, delta.visited, delta.straightLine, 0}
			if !isWithinBounds(m, p) || p.straightLine > 2 {
				continue
			}
			key1 := getKey(p, p.straightLine)
			p.pathValue = point.pathValue + m[p.x][p.y]
			if val, ok := globalVis[key1]; ok && val <= p.pathValue {
				continue
			}
			// p.visited = copyMap(point.visited)
			// p.visited = point.visited
			// queue = append(queue, p)
			heap.Push(&pq, &p)
		}

	}
	return minValue
}

func debugQueue(queue []Point) {
	for _, point := range queue {
		fmt.Print(point.pathValue, " ")
		// fmt.Print(fmt.Sprintf("%d,%d,%c ", point.x, point.y, point.dir))

	}
	fmt.Println()
}

// func solve(m [][]int, point Point, straightDirection int, visited map[string]int, visitedLocal map[string]bool) int {
// 	// if count == 500 {
// 	// 	panic("count exceeded")
// 	// }

// 	count++
// 	key := getKey(point, straightDirection)
// 	_, ok := visitedLocal[key]
// 	if ok {
// 		return 10000000000000000
// 	}
// 	visitedLocal[key] = true
// 	// fmt.Println(key, visited)
// 	// fmt.Println(key)
// 	if val, ok := visited[key]; ok {
// 		return val
// 	}
// 	if point.x == len(m)-1 && point.y == len(m[0])-1 {
// 		return m[point.x][point.y]
// 	}

// 	deltas := make([]Point, 0)
// 	if point.dir == 'L' {
// 		deltas = append(deltas, Point{0, -1, 'L'})
// 		deltas = append(deltas, Point{1, 0, 'D'})
// 		deltas = append(deltas, Point{-1, 0, 'U'})
// 	} else if point.dir == 'R' {
// 		deltas = append(deltas, Point{0, 1, 'R'})
// 		deltas = append(deltas, Point{1, 0, 'D'})
// 		deltas = append(deltas, Point{-1, 0, 'U'})
// 	} else if point.dir == 'U' {
// 		deltas = append(deltas, Point{0, -1, 'L'})
// 		deltas = append(deltas, Point{0, 1, 'R'})
// 		deltas = append(deltas, Point{-1, 0, 'U'})
// 	} else if point.dir == 'D' {
// 		deltas = append(deltas, Point{0, -1, 'L'})
// 		deltas = append(deltas, Point{0, 1, 'R'})
// 		deltas = append(deltas, Point{1, 0, 'D'})
// 	}

// 	minValue := 10000000000000000
// 	minKey := "X"
// 	for _, delta := range deltas {
// 		p := Point{point.x + delta.x, point.y + delta.y, delta.dir}
// 		// _, ok := visitedLocal[getKey(p, straightDirection)]
// 		if !isWithinBounds(m, p) || (delta.dir == point.dir && straightDirection == 3) {
// 			continue
// 		}
// 		if delta.dir == point.dir {
// 			mapCopy := copyMap(visitedLocal)
// 			val := solve(m, p, straightDirection+1, visited, mapCopy)
// 			fmt.Println(key, " -> ", getKey(p, straightDirection), " -> ", val)
// 			appendToFile(fmt.Sprintf("%s -> %s -> %d\n", key, getKey(p, straightDirection+1), val))
// 			if val < minValue {
// 				minKey = getKey(p, straightDirection)
// 				minValue = val
// 			}
// 		} else {
// 			mapCopy := copyMap(visitedLocal)
// 			val := solve(m, p, 0, visited, mapCopy)
// 			appendToFile(fmt.Sprintf("%s -> %s -> %d\n", key, getKey(p, 0), val))
// 			if val < minValue {
// 				minKey = getKey(p, straightDirection)
// 				minValue = val
// 			}
// 		}
// 	}
// 	visited[key] = minValue + m[point.x][point.y]
// 	fmt.Println(key, " -> ", minKey)
// 	appendToFile(fmt.Sprintf("%s -> %s\n", key, minKey))
// 	return minValue + m[point.x][point.y]
// }

func appendToFile(str string) {
	fd, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file")
		panic(err)
	}
	defer fd.Close()
	fd.WriteString(str)
}

func clearFile() {
	fd, err := os.OpenFile("output.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file")
		panic(err)
	}
	defer fd.Close()
}

func isWithinBounds(m [][]int, point Point) bool {
	return point.x >= 0 && point.x < len(m) && point.y >= 0 && point.y < len(m[0])
}

func getKey(point Point, straightDirection int) string {
	return fmt.Sprintf("%d,%d,%c,%d", point.x, point.y, point.dir, straightDirection)
}

func copyMap(visited *map[string]int) *map[string]int {
	newMap := make(map[string]int)
	for k, v := range *visited {
		newMap[k] = v
	}
	return &newMap
}

type PriorityQueue []*Point

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].pathValue < pq[j].pathValue
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Point)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
