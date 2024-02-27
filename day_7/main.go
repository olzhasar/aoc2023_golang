package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
Cards
A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2

Combinations
Five of a kind, where all five cards have the same label: AAAAA
Four of a kind, where four cards have the same label and one card has a different label: AA8AA
Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
High card, where all cards' labels are distinct: 23456
*/

func getCombinationScore(hand string) int {
	kinds := make(map[string]int)
	largest := 0

	for _, card := range hand {
		kinds[string(card)]++
		if kinds[string(card)] > largest {
			largest = kinds[string(card)]
		}
	}

	switch len(kinds) {
	case 1:
		return 6
	case 2:
		if largest == 4 {
			return 5
		}
		return 4
	case 3:
		if largest == 3 {
			return 3
		}
		return 2
	case 4:
		return 1
	case 5:
		return 0
	}

	panic("Invalid hand")
}

func getCardSore(card string) int {
	switch card {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 11
	case "T":
		return 10
	default:
		result, err := strconv.Atoi(card)
		if err != nil {
			panic(err)
		}
		return result
	}
}

func getScore(hand string) int {
	score := 0
	base := 16

	combScore := getCombinationScore(hand)

	pow := func(x, y int) int {
		return int(math.Pow(float64(x), float64(y)))
	}

	score += combScore * pow(base, 6)

	for i, card := range hand {
		score += getCardSore(string(card)) * pow(base, 5-i)
	}

	return score
}

func getBid(val string) int {
	result, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return result
}

func parseHand(hand string) [2]int {
	parts := strings.Fields(hand)
	score := getScore(parts[0])
	bid := getBid(parts[1])

	return [2]int{score, bid}
}

func GetTotalWinnings(input []string) int {
	hands := make([][2]int, 0)

	for _, line := range input {
		hands = append(hands, parseHand(line))
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i][0] < hands[j][0]
	})

	result := 0

	for i, hand := range hands {
		result += hand[1] * (i + 1)
	}

	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	result := GetTotalWinnings(input)

	fmt.Println(result)
}
