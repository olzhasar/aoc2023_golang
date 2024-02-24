package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getCardPoints(card string) int {
	card = strings.Split(card, ":")[1]

	divided := strings.Split(card, "|")

	winningNumbers := make(map[int]bool)

	points := 0

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
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}

	return points
}

func GetTotalPoints(input []string) int {
	total := 0

	for _, str := range input {
		total += getCardPoints(str)
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

	result := GetTotalPoints(input)

	fmt.Println(result)
}
