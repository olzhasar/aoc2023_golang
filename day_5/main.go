package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func parseNum(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

func getSeeds(input string) []int {
	result := make([]int, 0)

	numPart := strings.Split(input, ":")[1]
	numsStrings := strings.Fields(numPart)

	for _, num := range numsStrings {
		result = append(result, parseNum(num))
	}

	return result
}

func parseRow(input string) (int, int, int) {
	valueStrings := strings.Fields(input)

	var values [3]int

	for i, str := range valueStrings {
		values[i] = parseNum(str)
	}

	return values[0], values[1], values[2]
}

func processRow(input string, result []int, visited []bool) []int {
	destStart, sourceStart, rangeLen := parseRow(input)

	for i := 0; i < len(result); i++ {
		if visited[i] {
			continue
		}

		if result[i] >= sourceStart && result[i] < sourceStart+rangeLen {
			diff := result[i] - sourceStart
			result[i] = destStart + diff
			visited[i] = true
		}
	}

	return result
}

func resetVisited(visited []bool) {
	for i := range visited {
		visited[i] = false
	}
}

func processMap(input []string, result []int) []int {
	visited := make([]bool, len(result))
	resetVisited(visited)

	for _, row := range input {
		firstRune, _ := utf8.DecodeRuneInString(row)
		if unicode.IsDigit(firstRune) {
			result = processRow(row, result, visited)
		} else {
			resetVisited(visited)
		}
	}

	return result
}

func GetLowestLocation(input []string) int {
	seeds := getSeeds(input[0])

	result := processMap(input[1:], seeds)

	return slices.Min(result)
}

func GetLowestLocationPart2(input []string) int {
	return 0
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

	result := GetLowestLocation(input)

	fmt.Println(result)
}
