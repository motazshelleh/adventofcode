package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_BLUE  = 14
	MAX_RED   = 12
	MAX_GREEN = 13
)

func main() {
	fd, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		panic(err)
	}
	defer fd.Close()
	fileReader := bufio.NewScanner(fd)
	sum := 0
	for fileReader.Scan() {
		line := fileReader.Text()
		gameSplit := strings.Split(line, ":")
		gameId := strings.Split(gameSplit[0], " ")[1]
		blue, red, green := getCubesCount(gameSplit[1])
		if blue <= MAX_BLUE && red <= MAX_RED && green <= MAX_GREEN {
			n, err := strconv.Atoi(gameId)
			if err != nil {
				fmt.Println("Error converting string to int " + gameId)
				panic(err)
			}
			sum += n
		}
	}

	fmt.Println(sum)
}

func getCubesCount(line string) (int, int, int) {
	blue := 0
	red := 0
	green := 0

	for _, s := range strings.Split(line, ";") {
		for _, c := range strings.Split(s, ",") {
			parts := strings.Split(strings.TrimSpace(c), " ")
			switch parts[1] {
			case "blue":
				blue = findMax(parts[0], blue)
			case "red":
				red = findMax(parts[0], red)
			case "green":
				green = findMax(parts[0], green)
			}
		}
	}
	return blue, red, green
}

func findMax(a string, b int) int {
	n, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("Error converting string to int " + a)
		panic(err)
	}
	return int(math.Max(float64(n), float64(b)))
}
