package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type znode struct {
	steps    int
	zname    string
	insIndex int
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

	startingNodes := make([]string, 0)
	for k, _ := range nodeMap {
		if k[len(k)-1] == 'A' {
			startingNodes = append(startingNodes, k)
		}
	}
	fmt.Println(startingNodes)

	r := []rune(instructions)
	zNodesMap := make(map[string]znode)
	for k, _ := range nodeMap {
		if k[len(k)-1] != 'Z' {
			continue
		}
		for i := 0; i < len(instructions); i++ {
			steps, zNodename, index := stepsToZZZ(nodeMap[k], nodeMap, r, i)
			key := k + ":" + strconv.Itoa(i)

			zNodesMap[key] = znode{
				steps:    steps,
				zname:    zNodename,
				insIndex: index,
			}
		}
	}

	for _, k := range startingNodes {
		for i := 0; i < len(instructions); i++ {
			steps, zNodename, index := stepsToZZZ(nodeMap[k], nodeMap, r, i)
			key := k + ":" + strconv.Itoa(i)

			zNodesMap[key] = znode{
				steps:    steps,
				zname:    zNodename,
				insIndex: index,
			}
		}
	}

	values := make([]int, 0)
	for _, n := range startingNodes {
		key := n + ":0"
		values = append(values, zNodesMap[key].steps)
	}
	fmt.Println(values)

	indexes1 := []int{1, 1, 1, 1, 1, 1}
	for true {
		prevValue := values[0] * indexes1[0]
		allSame := true
		minValue := prevValue
		minIndex := 0
		maxValue := prevValue
		for i := 1; i < len(values); i++ {
			if values[i]*indexes1[i] != prevValue {
				allSame = false
			}
			if values[i]*indexes1[i] < minValue {
				minValue = values[i] * indexes1[i]
				minIndex = i
			}
			if values[i]*indexes1[i] > maxValue {
				maxValue = values[i] * indexes1[i]
			}
		}
		if allSame {
			fmt.Println(minValue)
			fmt.Println(indexes1)
			return
		}
		// fmt.Println(indexes1[minIndex]*values[minIndex], maxValue)
		for indexes1[minIndex]*values[minIndex] < maxValue {
			indexes1[minIndex]++
		}
		// fmt.Println(indexes1)
	}

	// // fmt.Println(zNodesMap)
	// numberOfNodesAtZ := 0
	// //from a snapshot of the map, find the minimum steps to ZZZ
	// startingNodesIndexes := []int{280, 280, 280, 280, 280, 280}
	// startingNodesSteps := []int{12927600766034, 12927600759722, 12927600761194, 12927600762783, 12927600763331, 12927600765714}
	// startingNodes = []string{"DLZ", "ZZZ", "RGZ", "BGZ", "HBZ", "NTZ"}
	// count := 0
	// for numberOfNodesAtZ < len(startingNodes) {

	// 	// find fine nodeIndex of the minimum steps
	// 	minSteps := startingNodesSteps[0]
	// 	maxSteps := startingNodesSteps[0]
	// 	minStepsIndex := 0
	// 	tempArr := []int{280, 280, 280, 280, 280, 280}
	// 	tempArr[0] = 12927600769609 - startingNodesSteps[0]
	// 	for i := 1; i < len(startingNodes); i++ {
	// 		if startingNodesSteps[i] < minSteps {
	// 			minSteps = startingNodesSteps[i]
	// 			minStepsIndex = i
	// 		}
	// 		if startingNodesSteps[i] > maxSteps {
	// 			maxSteps = startingNodesSteps[i]
	// 		}
	// 		tempArr[i] = 12927600769609 - startingNodesSteps[i]
	// 	}
	// 	fmt.Println(tempArr)

	// 	if maxSteps > 12927600769609 {
	// 		fmt.Println("excceeds 12927600769609")
	// 		break
	// 	}

	// 	// 12083 19951 17141 22199 16579 14893

	// 	startingNode := startingNodes[minStepsIndex]
	// 	nodeIndex := minStepsIndex
	// 	// for nodeIndex, startingNode := range startingNodes {
	// 	key := startingNode + ":" + strconv.Itoa(startingNodesIndexes[nodeIndex])
	// 	_, ok := zNodesMap[key]
	// 	if !ok {
	// 		fmt.Println("not ok", key)
	// 		break
	// 	}
	// 	startingNodes[nodeIndex] = zNodesMap[key].zname
	// 	startingNodesIndexes[nodeIndex] = zNodesMap[key].insIndex
	// 	startingNodesSteps[nodeIndex] += zNodesMap[key].steps
	// 	// }
	// 	allNodesHaveSameSteps := true
	// 	for i := 1; i < len(startingNodes); i++ {
	// 		if startingNodesSteps[i] != startingNodesSteps[i-1] {
	// 			allNodesHaveSameSteps = false
	// 			break
	// 		}
	// 	}
	// 	// fmt.Println(startingNodesSteps)
	// 	if count%1 == 0 {
	// 		fmt.Println(startingNodesSteps, startingNode, startingNodesIndexes, startingNodes)
	// 	}
	// 	if allNodesHaveSameSteps {
	// 		fmt.Println(startingNodesSteps)
	// 		break
	// 	}
	// 	count++
	// }

	// 	steps := 0
	// outerLoopLabel:
	// 	for true {
	// 		for _, c := range instructions {
	// 			steps++
	// 			numberOfNodesAtZ := 0
	// 			for i, n := range startingNodes {
	// 				if c == 'L' {
	// 					startingNodes[i] = nodeMap[n].left
	// 				} else {
	// 					startingNodes[i] = nodeMap[n].right
	// 				}
	// 				if startingNodes[i][len(startingNodes[i])-1] == 'Z' {
	// 					numberOfNodesAtZ += 1
	// 				}
	// 			}
	// 			if numberOfNodesAtZ > 3 {
	// 				fmt.Println(numberOfNodesAtZ, "steps so far: ", steps)
	// 			}
	// 			if numberOfNodesAtZ == len(startingNodes) {
	// 				fmt.Println("steps: ", steps)
	// 				break outerLoopLabel
	// 			}
	// 		}
	// 	}

	// fmt.Println(instructions, nodeMap)
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

func stepsToZZZ(currentNode node, nodeMap map[string]node, instructions []rune, i int) (int, string, int) {
	steps := 0
	for true {
		for ; i < len(instructions); i++ {
			c := instructions[i]
			if c == 'L' {
				currentNode = nodeMap[currentNode.left]
			} else {
				currentNode = nodeMap[currentNode.right]
			}
			steps++
			// fmt.Println(currentNode.name, i)
			if currentNode.name[len(currentNode.name)-1] == 'Z' {
				return steps, currentNode.name, i
			}
		}
		i = 0
	}

	return steps, "", -1
}
