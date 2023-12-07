package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards    []rune
	bid      int
	strength int
}

func main() {
	fd, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		panic(err)
	}
	defer fd.Close()

	fileReader := bufio.NewScanner(fd)
	arr := make([]hand, 0)
	for fileReader.Scan() {
		line := fileReader.Text()
		arr = append(arr, parseStringLineIntoArray(line))
	}

	sort.Slice(arr, func(i, j int) bool {
		return compareHands(arr[i], arr[j]) < 0
	})

	ans := 0
	for i, h := range arr {
		ans += (h.bid * (i + 1))
	}

	// fmt.Println(arr)
	fmt.Println(ans)
}

func parseStringLineIntoArray(line string) hand {
	parts := strings.Split(strings.TrimSpace(line), " ")
	num, _ := strconv.Atoi(parts[1])
	card := getRuneArrFromString(parts[0])
	h := hand{
		cards:    card,
		bid:      num,
		strength: getHandStrength(card),
	}
	return h
}

func getRuneArrFromString(str string) []rune {
	return []rune(str)
}

// higher strength is better hand
func getHandStrength(card []rune) int {
	cardsMap := make(map[rune]int)
	jCount := 0
	max := 0
	for _, c := range card {
		if c == 'J' {
			jCount++
		} else {
			cardsMap[c]++
			if cardsMap[c] > max {
				max = cardsMap[c]
			}
		}
	}

	if jCount == 5 {
		// hand has 5 jokers
		return 7
	}

	for k, v := range cardsMap {
		if v == max {
			// add jokers to the max card
			cardsMap[k] += jCount
			break
		}
	}
	if len(cardsMap) == 1 {
		// all cards are the same
		return 7
	}
	if len(cardsMap) == 2 {
		// hand has either 4 of a kind (AA8AA) or full house (23332)
		for _, v := range cardsMap {
			if v == 3 {
				return 5
			}
		}
		return 6
	}
	if len(cardsMap) == 3 {
		// hand has either 3 of a kind (AAA98) or 2 pair (AABBK)
		for _, v := range cardsMap {
			if v == 3 {
				return 4
			}
		}
		return 3
	}
	// hand has either 1 pair (AAK98) or high card (AKJ98)
	return 6 - len(cardsMap)
}

func compareHands(h1 hand, h2 hand) int {
	if h1.strength > h2.strength {
		return 1
	}
	if h1.strength < h2.strength {
		return -1
	}
	for i := 0; i < len(h1.cards); i++ {
		h1CardStrength := getCardStrength(h1.cards[i])
		h2CardStrength := getCardStrength(h2.cards[i])
		if h1CardStrength > h2CardStrength {
			return 1
		}
		if h1CardStrength < h2CardStrength {
			return -1
		}
	}
	return 0
}

func getCardStrength(card rune) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'T':
		return 10
	case 'J':
		return 1
	default:
		return int(card - '0')
	}
}
