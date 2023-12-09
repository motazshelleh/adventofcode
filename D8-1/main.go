package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type hand struct {
	cards    []rune
	bid      int
	strength int
}

type node struct {
	name  string
	left  string
	right string
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
	instructions := fileReader.Text()
	fileReader.Scan()
	fileReader.Text()

	nodeMap := make(map[string]node)
	for fileReader.Scan() {
		line := fileReader.Text()
		parseStringLine(line, &nodeMap)
	}

	currentNode := nodeMap["AAA"]
	steps := 0
outerLoopLabel:
	for true {
		for _, c := range instructions {
			if c == 'L' {
				currentNode = nodeMap[currentNode.left]
			} else {
				currentNode = nodeMap[currentNode.right]
			}
			steps++
			if currentNode.name == "ZZZ" {
				fmt.Println("steps: ", steps)
				break outerLoopLabel
			}
		}
	}

	fmt.Println(instructions, nodeMap)
}

func parseStringLine(line string, m *map[string]node) {
	parts := strings.Split(line, " = ")
	n := node{}
	n.name = parts[0]
	parts = strings.Split(parts[1][1:len(parts[1])-1], ", ")
	n.left = parts[0]
	n.right = parts[1]
	(*m)[n.name] = n
}
