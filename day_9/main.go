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

func GetExtrapolatedSum(input []string) int {
	total := 0

	for _, row := range input {
		nums := parseRow(row)
		total += predictNextValue(nums)
	}

	return total
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

	result := GetExtrapolatedSum(input)

	fmt.Println(result)
}
