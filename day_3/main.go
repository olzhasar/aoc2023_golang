package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type Part struct {
	x_start int
	y_start int
	y       int
	number  int
}

func getSymbolMatrix(input []string) [][]bool {
	length := len(input)

	matrix := make([][]bool, length)

	for i, row := range input {
		for _, char := range row {
			if char == '.' || unicode.IsDigit(char) {
				matrix[i] = append(matrix[i], false)
			} else {
				matrix[i] = append(matrix[i], true)
			}
		}
	}

	return matrix
}

func getAllParts(input []string) []Part {
	result := []Part{}

	start := -1
	end := -1

	appendPart := func(row string, rowNumber int) {
		number := 0

		for i := start; i <= end; i++ {
			number = number*10 + int(row[i]-'0')
		}

		result = append(result, Part{start, end, rowNumber, number})

		start = -1
		end = -1
	}

	for i, row := range input {
		for j, char := range row {
			if unicode.IsDigit(char) {
				if start == -1 {
					start = j
				}
				end = j
			} else {
				if start != -1 {
					appendPart(row, i)
				}
			}
		}

		if start != -1 {
			appendPart(row, i)
		}
	}

	return result
}

func hasAdjacentSymbol(matrix [][]bool, part Part, m int, n int) bool {
	directions := [8][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}

	for i := part.x_start; i <= part.y_start; i++ {
		for _, direction := range directions {
			x := direction[0] + i
			y := direction[1] + part.y

			if x >= 0 && x < m && y >= 0 && y < n && matrix[y][x] {
				return true
			}
		}
	}

	return false
}

func getPartNumbersSum(input []string) int {
	M := len(input)
	N := len(input[0])

	symbolsMatrix := getSymbolMatrix(input)

	parts := getAllParts(input)

	total := 0

	for _, part := range parts {
		if hasAdjacentSymbol(symbolsMatrix, part, M, N) {
			total += part.number
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

	result := getPartNumbersSum(input)

	fmt.Println(result)
}
