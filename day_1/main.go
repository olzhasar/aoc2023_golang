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
	return 0
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
		result += calculateCalibration(scanner.Text())
	}

	fmt.Println(result)
}
