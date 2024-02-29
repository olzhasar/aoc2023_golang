package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return result
}

func parseRow(row string) []int {
	var nums []int

	for _, numStr := range strings.Fields(row) {
		nums = append(nums, parseInt(numStr))
	}

	return nums
}

func predictNextValue(nums []int) int {
	length := len(nums)

	if length == 0 {
		return 0
	}

	if length == 1 {
		return nums[0]
	}

	var diffs []int

	for i := 1; i < length; i++ {
		diffs = append(diffs, nums[i]-nums[i-1])
	}

	return nums[length-1] + predictNextValue(diffs)
}

func predictPrevValue(nums []int) int {
	length := len(nums)

	if length == 0 {
		return 0
	}

	if length == 1 {
		return nums[0]
	}

	var diffs []int

	for i := 1; i < length; i++ {
		diffs = append(diffs, nums[i-1]-nums[i])
	}

	return nums[0] + predictPrevValue(diffs)
}

func getSum(input []string, predictFunc func([]int) int) int {
	total := 0

	for _, row := range input {
		nums := parseRow(row)
		total += predictFunc(nums)
	}

	return total
}

func GetExtrapolatedSum(input []string) int {
	return getSum(input, predictNextValue)
}

func GetExtrapolatedSumPartTwo(input []string) int {
	return getSum(input, predictPrevValue)
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

	result := GetExtrapolatedSumPartTwo(input)

	fmt.Println(result)
}
