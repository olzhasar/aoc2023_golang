package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func getSeeds(input string) []int {
	result := make([]int, 0)

	numPart := strings.Split(input, ":")[1]
	numsStrings := strings.Fields(numPart)

	for _, num := range numsStrings {
		n, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}

		result = append(result, n)
	}

	return result
}

func processRow(input string, matrix [][]int, iteration int) {
	valueStrings := strings.Fields(input)

	if len(valueStrings) != 3 {
		panic("Invalid row")
	}

	parseNum := func(s string) int {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return num
	}

	destStart := parseNum(valueStrings[0])
	sourceStart := parseNum(valueStrings[1])
	rangeLen := parseNum(valueStrings[2])

	for i := range matrix {
		if len(matrix[i]) > iteration {
			continue
		}

		val := matrix[i][iteration-1]

		if val >= sourceStart && val < sourceStart+rangeLen {
			diff := val - sourceStart
			matrix[i] = append(matrix[i], destStart+diff)
		}
	}
}

func fillMissingValues(matrix [][]int, iteration int) {
	for i := range matrix {
		length := len(matrix[i])
		if length <= iteration {
			matrix[i] = append(matrix[i], matrix[i][length-1])
		}
	}
}

func processMap(input []string, matrix [][]int) {
	iteration := 0

	for _, row := range input {
		firstRune, _ := utf8.DecodeRuneInString(row)
		if unicode.IsDigit(firstRune) {
			processRow(row, matrix, iteration)
		} else if unicode.IsLetter(firstRune) {
			if iteration > 0 {
				fillMissingValues(matrix, iteration)
			}
			iteration++
		}
	}

	fillMissingValues(matrix, iteration)

}

func GetLowestLocation(input []string) int {
	seeds := getSeeds(input[0])
	var matrix [][]int

	for i := 0; i < len(seeds); i++ {
		matrix = append(matrix, []int{seeds[i]})
	}

	processMap(input[1:], matrix)

	lowest := matrix[0][len(matrix[0])-1]

	for i := range matrix[1:] {
		length := len(matrix[i+1])
		if matrix[i+1][length-1] < lowest {
			lowest = matrix[i+1][length-1]
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

	result := GetLowestLocation(input)

	fmt.Println(result)

}
