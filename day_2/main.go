package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Totals [3]int // RGB

func splitGameId(input string) (int, string) {
	splitted := strings.Split(input, ":")

	gameIdString := strings.TrimSpace(splitted[0])

	gameId, err := strconv.Atoi(gameIdString[5:])
	if err != nil {
		panic(err)
	}

	return gameId, splitted[1]
}

func getSetTotals(setString string) Totals {
	totals := Totals{0, 0, 0}

	colors := strings.Split(setString, ",")

	for _, color := range colors {
		splitted := strings.Split(strings.TrimSpace(color), " ")

		quantity, err := strconv.Atoi(splitted[0])
		if err != nil {
			panic(err)
		}

		switch splitted[1] {
		case "red":
			totals[0] += quantity
		case "green":
			totals[1] += quantity
		case "blue":
			totals[2] += quantity
		}
	}

	return totals
}

func parseTotals(input string) []Totals {
	result := []Totals{}
	setStrings := strings.Split(input, ";")

	for _, setString := range setStrings {
		totals := getSetTotals(setString)
		result = append(result, totals)

	}

	return result
}

func checkSet(setTotals Totals, totals Totals) bool {
	for i := 0; i < 3; i++ {
		if setTotals[i] > totals[i] {
			return false
		}
	}

	return true
}

func checkGame(input string, totals Totals) int {
	gameId, setsString := splitGameId(input)

	totalSets := parseTotals(setsString)

	for _, set := range totalSets {
		if !checkSet(set, totals) {
			return 0
		}
	}

	return gameId
}

func calcGamePower(input string) int {
	_, setsString := splitGameId(input)

	fewest := []int{0, 0, 0}

	totalSets := parseTotals(setsString)

	for _, set := range totalSets {
		for i := 0; i < 3; i++ {
			if set[i] > fewest[i] {
				fewest[i] = set[i]
			}
		}
	}

	result := 1

	for i := 0; i < 3; i++ {
		result *= fewest[i]
	}

	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	result := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result += calcGamePower(scanner.Text())
	}

	fmt.Println(result)
}
