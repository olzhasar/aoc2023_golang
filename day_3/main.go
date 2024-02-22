package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var DIRS = [8][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
	{-1, -1},
	{-1, 1},
	{1, -1},
	{1, 1},
}

func getMatrix(input []string) (matrix [][]rune, m int, n int) {
	m = len(input)

	matrix = make([][]rune, m)

	for i, row := range input {
		matrix[i] = []rune(row)
	}

	n = len(matrix[0])

	return matrix, m, n
}

func isSymbol(char rune) bool {
	return char != '.' && !unicode.IsDigit(char)
}

func isDigit(char rune) bool {
	return unicode.IsDigit(char)
}

func isStar(char rune) bool {
	return char == '*'
}

func hasAdjacentSymbol(matrix [][]rune, m int, n int, row int, column int) bool {
	for _, direction := range DIRS {
		r := direction[0] + row
		c := direction[1] + column

		if r >= 0 && r < m && c >= 0 && c < n && isSymbol(matrix[r][c]) {
			return true
		}
	}
	return false
}

func getPartNumbersSum(input []string) int {
	matrix, m, n := getMatrix(input)

	fmt.Println("m:", m, "n:", n)

	total := 0

	current := -1
	has_adj := false

	for i, row := range matrix {
		for j, char := range row {
			if isDigit(char) {
				val := int(char - '0')
				if current == -1 {
					current = val
				} else {
					current = current*10 + val
				}
				has_adj = has_adj || hasAdjacentSymbol(matrix, m, n, i, j)
			} else {
				if current != -1 && has_adj {
					total += current
				}
				current = -1
				has_adj = false
			}
		}
		if current != -1 && has_adj {
			total += current
		}
		current = -1
		has_adj = false
	}

	return total
}

func getAllNumbers(matrix [][]rune) (numbersMatrix [][]int, numbers []int) {
	numbersMatrix = make([][]int, len(matrix))
	numbers = make([]int, 0)

	current := -1
	index := 0

	for i, row := range matrix {
		for _, char := range row {
			if isDigit(char) {
				val := int(char - '0')
				if current == -1 {
					current = val
				} else {
					current = current*10 + val
				}
				numbersMatrix[i] = append(numbersMatrix[i], index)
			} else {
				if current != -1 {
					numbers = append(numbers, current)
					index++
				}
				current = -1
				numbersMatrix[i] = append(numbersMatrix[i], -1)
			}
		}
		if current != -1 {
			numbers = append(numbers, current)
			index++
		}
		current = -1
	}

	return numbersMatrix, numbers
}

func calcGearRatio(numbersMatrix [][]int, numbers []int, m int, n int, row int, col int) int {
	ratio := 1
	metNumbers := make([]int, 0)

	contains := func(val int) bool {
		for _, num := range metNumbers {
			if num == val {
				return true
			}
		}
		return false
	}

	for _, direction := range DIRS {
		r := row + direction[0]
		c := col + direction[1]
		index := numbersMatrix[r][c]
		if r >= 0 && r < m && c >= 0 && c < n && index != -1 {
			number := numbers[index]

			if !contains(number) {
				if len(metNumbers) > 1 {
					return 0
				}
				metNumbers = append(metNumbers, number)
				ratio *= number
			}
		}
	}

	if len(metNumbers) != 2 {
		return 0
	}

	return ratio
}

func getGearRatioSum(input []string) int {
	matrix, m, n := getMatrix(input)

	numbersMatrix, numbers := getAllNumbers(matrix)

	total := 0

	for i, row := range matrix {
		for j, char := range row {
			if isStar(char) {
				total += calcGearRatio(numbersMatrix, numbers, m, n, i, j)
			}
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

	result := getGearRatioSum(input)

	fmt.Println(result)
}
