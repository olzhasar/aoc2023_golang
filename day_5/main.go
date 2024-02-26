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

func parseRow(input string) (int, int, int) {
	valueStrings := strings.Fields(input)

	var values [3]int

	for i, str := range valueStrings {
		values[i] = parseNum(str)
	}

	return values[0], values[1], values[2]
}

func getSeeds(input string) []int {
	var result []int

	numPart := strings.Split(input, ":")[1]
	numsStrings := strings.Fields(numPart)

	for _, num := range numsStrings {
		result = append(result, parseNum(num))
	}

	return result
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

func GetLowestLocation(input []string) int {
	result := getSeeds(input[0])

	visited := make([]bool, len(result))

	resetVisited := func() {
		for i := range visited {
			visited[i] = false
		}
	}

	resetVisited()

	for _, row := range input {
		firstRune, _ := utf8.DecodeRuneInString(row)
		if unicode.IsDigit(firstRune) {
			result = processRow(row, result, visited)
		} else {
			resetVisited()
		}
	}

	return slices.Min(result)
}

func getSeedRanges(input string) [][2]int {
	var result [][2]int

	nums := getSeeds(input)

	for i := 0; i < len(nums)/2; i++ {
		start := nums[i*2]
		end := start + nums[i*2+1] - 1 // inclusive

		result = append(result, [2]int{start, end})
	}

	return result
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func processRowPart2(input string, initial [][2]int, result [][2]int) ([][2]int, [][2]int) {
	destStart, sourceStart, rangeLen := parseRow(input)
	sourceEnd := sourceStart + rangeLen - 1

	var newInitial [][2]int

	for _, row := range initial {
		start := row[0]
		end := row[1]

		overlapStart := max(start, sourceStart)
		overlapEnd := min(end, sourceEnd)

		if overlapEnd >= overlapStart {
			diff := overlapStart - sourceStart

			result = append(result, [2]int{destStart + diff, destStart + diff + overlapEnd - overlapStart})
		}

		leftStart := start
		leftEnd := min(sourceStart-1, end)

		if leftEnd >= leftStart {
			newInitial = append(newInitial, [2]int{leftStart, leftEnd})
		}

		rightStart := max(sourceEnd+1, start)
		rightEnd := end

		if rightEnd >= rightStart {
			newInitial = append(newInitial, [2]int{rightStart, rightEnd})
		}
	}

	return newInitial, result
}

func GetLowestLocationPart2(input []string) int {
	initial := getSeedRanges(input[0])

	var result [][2]int

	for _, row := range input[1:] {
		firstRune, _ := utf8.DecodeRuneInString(row)
		if unicode.IsDigit(firstRune) {
			initial, result = processRowPart2(row, initial, result)
		} else if unicode.IsLetter(firstRune) {
			for _, result_row := range result {
				initial = append(initial, result_row)
			}
			result = make([][2]int, 0)
		}
	}

	for _, row := range initial {
		result = append(result, row)
	}

	lowest := result[0][0]
	for _, row := range result {
		if row[0] < lowest {
			lowest = row[0]
		}
	}

	return lowest
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

	result := GetLowestLocationPart2(input)

	fmt.Println(result)
}
