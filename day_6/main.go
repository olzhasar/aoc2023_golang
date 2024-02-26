package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseNumber(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

func parseNumbers(input string) []int {
	var result []int

	numsString := strings.Split(input, ":")[1]
	for _, timeStr := range strings.Fields(numsString) {
		result = append(result, parseNumber(timeStr))
	}

	return result
}

func parseSingleNumber(input string) int {
	numString := strings.Split(input, ":")[1]
	cleaned := strings.ReplaceAll(numString, " ", "")
	return parseNumber(cleaned)
}

func getNumWaysToWin(time int, record int) int {
	result := 0

	for i := 1; i <= time; i++ {
		dist := i * (time - i)
		if dist > record {
			result++
		} else if result > 0 {
			break
		}
	}

	return result
}

func GetNumWaysToWinMultiplied(input []string) int {
	timeArr := parseNumbers(input[0])
	distanceArr := parseNumbers(input[1])

	result := 1

	for i := 0; i < len(timeArr); i++ {
		result *= getNumWaysToWin(timeArr[i], distanceArr[i])
	}

	return result
}

func GetNumWaysToWinMultipliedPartTwo(input []string) int {
	time := parseSingleNumber(input[0])
	distance := parseSingleNumber(input[1])

	return getNumWaysToWin(time, distance)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	result := GetNumWaysToWinMultipliedPartTwo(input)

	fmt.Println(result)
}
