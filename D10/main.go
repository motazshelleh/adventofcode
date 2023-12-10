package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x, y, steps int
}

type direction struct {
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

	mapping := map[rune]map[direction][]rune{
		'+': {
			{-1, 0}: {'|', '7', 'F'},
			{1, 0}:  {'|', 'J', 'L'},
			{0, 1}:  {'-', 'J', '7'},
			{0, -1}: {'-', 'F', 'L'},
		},
		'|': {
			{-1, 0}: {'|', '7', 'F'},
			{1, 0}:  {'|', 'J', 'L'},
		},
		'-': {
			{0, 1}:  {'-', 'J', '7'},
			{0, -1}: {'-', 'F', 'L'},
		},
		'J': {
			{-1, 0}: {'|', '7', 'F'},
			{0, -1}: {'-', 'F', 'L'},
		},
		'L': {
			{-1, 0}: {'|', '7', 'F'},
			{0, 1}:  {'-', 'J', '7'},
		},
		'F': {
			{1, 0}: {'|', 'J', 'L'},
			{0, 1}: {'-', 'J', '7'},
		},
		'7': {
			{1, 0}:  {'|', 'J', 'L'},
			{0, -1}: {'-', 'F', 'L'},
		},
	}

	sketch := make([][]rune, 0)
	for fileReader.Scan() {
		line := fileReader.Text()
		sketch = append(sketch, []rune(line))
	}

	startX, startY := findStartingPoint(sketch)
	sketch[startX][startY] = '+'

	flood(sketch, startX, startY, mapping)

	sketch[startX][startY] = '-' // manually set to - (changed the shape of the start pipe) TODO: fix this
	visited := walkConnected(sketch, startX, startY)
	walkWall(sketch, startX, startY, visited)
	fmt.Println(visited[104][37])
}

func findStartingPoint(sketch [][]rune) (int, int) {
	for i, row := range sketch {
		for j, cell := range row {
			if cell == 'S' {
				return i, j
			}
		}
	}
	return -1, -1
}

func flood(sketch [][]rune, x, y int, mapping map[rune]map[direction][]rune) [][]bool {
	visited := make([][]bool, len(sketch))
	for i := range visited {
		visited[i] = make([]bool, len(sketch[0]))
	}
	q := make([]point, 0)
	q = append(q, point{x, y, 0})
	maxSteps := 0
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if visited[cur.x][cur.y] {
			continue
		}
		visited[cur.x][cur.y] = true
		allowed := getAllowedDirections(cur, sketch, visited, mapping)
		if cur.x == x && cur.y == y {
			fmt.Println("allowed for S", allowed)
		}
		for _, p := range allowed {
			q = append(q, p)
			if p.steps > maxSteps {
				maxSteps = p.steps
			}
		}
	}
	fmt.Println(maxSteps)
	return visited
}

func getAllowedDirections(p point, sketch [][]rune, visited [][]bool, mapping map[rune]map[direction][]rune) []point {
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	x := p.x
	y := p.y
	allowed := make([]point, 0)
	for _, dir := range directions {
		newX, newY := x+dir[0], y+dir[1]
		if newX < 0 || newX >= len(sketch) || newY < 0 || newY >= len(sketch[0]) || visited[newX][newY] || sketch[newX][newY] == '.' {
			continue
		}

		allowedPipes := mapping[sketch[x][y]][direction{dir[0], dir[1]}]
		for _, pipe := range allowedPipes {
			if pipe == sketch[newX][newY] {
				allowed = append(allowed, point{newX, newY, p.steps + 1})
				break
			}
		}
	}
	return allowed
}

func walkWall(sketch [][]rune, originalX int, originalY int, originalVisited [][]bool) {
	visited := make([][]bool, len(sketch))
	for i := range visited {
		visited[i] = make([]bool, len(sketch[0]))
	}
	x := originalX
	y := originalY
	total := 0
	currentDirection := direction{0, 1}
	insideDirection := 'U'
	mapDrawn := make([][]rune, len(sketch))
	for i := range mapDrawn {
		mapDrawn[i] = make([]rune, len(sketch[0]))
		for j := range mapDrawn[i] {
			mapDrawn[i][j] = '.'
		}
	}
	for {
		mapDrawn[x][y] = insideDirection
		insideDirection1, deltaX, deltaY := mapInsideDirection(sketch[x][y], insideDirection)
		insideDirection = insideDirection1
		switch sketch[x][y] {
		case '|':
			if isWithinBounds(sketch, x+deltaX, y+deltaY) && !originalVisited[x+deltaX][y+deltaY] {
				total += floodInside(sketch, x+deltaX, y+deltaY, visited, originalVisited, mapDrawn)
			}

			x += currentDirection.x
		case '-':
			if isWithinBounds(sketch, x+deltaX, y+deltaY) && !originalVisited[x+deltaX][y+deltaY] {
				total += floodInside(sketch, x+deltaX, y+deltaY, visited, originalVisited, mapDrawn)
			}
			y += currentDirection.y
		case 'J':
			if isWithinBounds(sketch, x+deltaX, y+deltaY) && !originalVisited[x+deltaX][y+deltaY] {
				total += floodInside(sketch, x+deltaX, y+deltaY, visited, originalVisited, mapDrawn)
			}

			if currentDirection.y == 1 {
				x -= 1
				currentDirection = direction{-1, 0}
			} else {
				y -= 1
				currentDirection = direction{0, -1}
			}
		case 'L':
			if isWithinBounds(sketch, x+deltaX, y+deltaY) && !originalVisited[x+deltaX][y+deltaY] {
				total += floodInside(sketch, x+deltaX, y+deltaY, visited, originalVisited, mapDrawn)
			}

			if insideDirection == 'L' {
				deltaX := 1
				deltaY := 0
				if isWithinBounds(sketch, x+deltaX, y+deltaY) && !originalVisited[x+deltaX][y+deltaY] {
					total += floodInside(sketch, x+deltaX, y+deltaY, visited, originalVisited, mapDrawn)
				}
			}

			if currentDirection.y == -1 {
				x -= 1
				currentDirection = direction{-1, 0}
			} else {
				y += 1
				currentDirection = direction{0, 1}
			}
		case 'F':
			if isWithinBounds(sketch, x+deltaX, y+deltaY) && !originalVisited[x+deltaX][y+deltaY] {
				total += floodInside(sketch, x+deltaX, y+deltaY, visited, originalVisited, mapDrawn)
			}

			if currentDirection.y == -1 {
				x += 1
				currentDirection = direction{1, 0}
			} else {
				y += 1
				currentDirection = direction{0, 1}
			}
		case '7':
			if isWithinBounds(sketch, x+deltaX, y+deltaY) && !originalVisited[x+deltaX][y+deltaY] {
				total += floodInside(sketch, x+deltaX, y+deltaY, visited, originalVisited, mapDrawn)
			}

			if currentDirection.y == 1 {
				x += 1
				currentDirection = direction{1, 0}
			} else {
				y -= 1
				currentDirection = direction{0, -1}
			}
		}

		if x == originalX && y == originalY {
			break
		}
	}
	fmt.Println("total areas inside: ", total)
}

func mapInsideDirection(currentPipe rune, insideDirection rune) (rune, int, int) {
	if currentPipe == '7' || currentPipe == 'L' {
		if insideDirection == 'L' {
			insideDirection = 'D'
		} else if insideDirection == 'R' {
			insideDirection = 'U'
		} else if insideDirection == 'U' {
			insideDirection = 'R'
		} else if insideDirection == 'D' {
			insideDirection = 'L'
		}
	} else if currentPipe == 'F' || currentPipe == 'J' {
		if insideDirection == 'L' {
			insideDirection = 'U'
		} else if insideDirection == 'R' {
			insideDirection = 'D'
		} else if insideDirection == 'U' {
			insideDirection = 'L'
		} else if insideDirection == 'D' {
			insideDirection = 'R'
		}
	}
	deltaX := 0
	deltaY := 0

	if currentPipe == '7' || currentPipe == 'F' {
		if insideDirection == 'L' {
			deltaY = -1
		} else if insideDirection == 'R' {
			deltaY = 1
		} else if insideDirection == 'U' {
			deltaX = -1
		} else if insideDirection == 'D' {
			deltaX = 1
		}
	} else if currentPipe == 'J' || currentPipe == 'L' {
		if insideDirection == 'L' {
			deltaY = -1
		} else if insideDirection == 'R' {
			deltaY = 1
		} else if insideDirection == 'U' {
			deltaX = -1
		} else if insideDirection == 'D' {
			deltaX = 1
		}
	} else if currentPipe == '|' {
		if insideDirection == 'L' {
			deltaY = -1
		} else if insideDirection == 'R' {
			deltaY = 1
		}
	} else if currentPipe == '-' {
		if insideDirection == 'U' {
			deltaX = -1
		} else if insideDirection == 'D' {
			deltaX = 1
		}
	}
	return insideDirection, deltaX, deltaY
}

func isWithinBounds(sketch [][]rune, x, y int) bool {
	return x >= 0 && x < len(sketch) && y >= 0 && y < len(sketch[0])
}

func floodInside(sketch [][]rune, x, y int, visited [][]bool, originalVisited [][]bool, drawingMap [][]rune) int {
	if visited[x][y] || originalVisited[x][y] {
		return 0
	}
	q := make([]point, 0)
	q = append(q, point{x, y, 0})
	total := 0
	visited[x][y] = true
	sketch[x][y] = 'I'
	drawingMap[x][y] = 'I'
	total++
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for _, dir := range directions {
			newX, newY := cur.x+dir[0], cur.y+dir[1]
			if newX < 0 || newX >= len(sketch) || newY < 0 || newY >= len(sketch[0]) || originalVisited[newX][newY] || visited[newX][newY] {
				continue
			}
			visited[newX][newY] = true
			sketch[newX][newY] = 'I'
			drawingMap[newX][newY] = 'I'
			q = append(q, point{newX, newY, cur.steps + 1})
			total++
		}
	}
	return total
}

func walkConnected(sketch [][]rune, originalX int, originalY int) [][]bool {
	visited := make([][]bool, len(sketch))
	for i := range visited {
		visited[i] = make([]bool, len(sketch[0]))
	}
	x := originalX
	y := originalY
	// currentDirection & insideDirection are set manually to match the starting pipe (depends on the input) TODO: fix this
	currentDirection := direction{0, 1}
	insideDirection := 'U'
	for {
		visited[x][y] = true
		insideDirection1, _, _ := mapInsideDirection(sketch[x][y], insideDirection)
		insideDirection = insideDirection1
		switch sketch[x][y] {
		case '|':
			x += currentDirection.x
		case '-':
			y += currentDirection.y
		case 'J':
			if currentDirection.y == 1 {
				x -= 1
				currentDirection = direction{-1, 0}
			} else {
				y -= 1
				currentDirection = direction{0, -1}
			}
		case 'L':
			if currentDirection.y == -1 {
				x -= 1
				currentDirection = direction{-1, 0}
			} else {
				y += 1
				currentDirection = direction{0, 1}
			}
		case 'F':
			if currentDirection.y == -1 {
				x += 1
				currentDirection = direction{1, 0}
			} else {
				y += 1
				currentDirection = direction{0, 1}
			}
		case '7':
			if currentDirection.y == 1 {
				x += 1
				currentDirection = direction{1, 0}
			} else {
				y -= 1
				currentDirection = direction{0, -1}
			}
		}
		if x == originalX && y == originalY {
			break
		}
	}
	return visited
}
