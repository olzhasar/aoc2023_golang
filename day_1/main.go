package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func calculateCalibration(str string) int {
	first := -1
	last := -1

	for _, char := range str {
		if unicode.IsDigit(char) {
			if first == -1 {
				first = int(char - '0')
			}
			last = int(char - '0')
		}
	}

	return first*10 + last
}

func calculateCalibrationPartTwo(str string) int {
	numberRunes := [][]rune{
		[]rune("one"),
		[]rune("two"),
		[]rune("three"),
		[]rune("four"),
		[]rune("five"),
		[]rune("six"),
		[]rune("seven"),
		[]rune("eight"),
		[]rune("nine"),
	}

	length := len(str)

	runes := []rune(str)

	getCurrentNumber := func(index int) int {
		r := runes[index]
		if unicode.IsDigit(r) {
			return int(r - '0')
		}

		checkNumber := func(i int, number []rune) bool {
			for _, r := range number {
				if i >= length || runes[i] != r {
					return false
				}
				i++
			}
			return true
		}

		for i, number := range numberRunes {
			if checkNumber(index, number) {
				return i + 1
			}
		}
		return -1
	}

	first := -1
	last := -1

	for i := 0; i < len(str); i++ {
		if number := getCurrentNumber(i); number != -1 {
			if first == -1 {
				first = number
			}
			last = number
		}
	}

	return first*10 + last
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
		result += calculateCalibrationPartTwo(scanner.Text())
	}

	fmt.Println(result)
}
