package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMatchingNumbers(card string) int {
	card = strings.Split(card, ":")[1]

	divided := strings.Split(card, "|")

	winningNumbers := make(map[int]bool)

	total := 0

	for _, str := range strings.Fields(divided[0]) {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		winningNumbers[num] = true
	}

	for _, str := range strings.Fields(divided[1]) {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		if winningNumbers[num] == true {
			total++
		}
	}

	return total
}

func getCardPoints(card string) int {
	matching := getMatchingNumbers(card)

	if matching == 0 {
		return 0
	}

	return 1 << (matching - 1)
}

func GetTotalPoints(input []string) int {
	total := 0

	for _, str := range input {
		total += getCardPoints(str)
	}
	return total
}

func GetCardsCount(input []string) int {
	total := 0

	var copies []int

	for i := 0; i < len(input); i++ {
		copies = append(copies, 1)
	}

	for i, str := range input {
		current := getMatchingNumbers(str)
		total += copies[i]

		for j := i + 1; j < i+current+1; j++ {
			copies[j] += copies[i]
		}
	}
	return total
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := []string{}

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	result := GetCardsCount(input)

	fmt.Println(result)
}
