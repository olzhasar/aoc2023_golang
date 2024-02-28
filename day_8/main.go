package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseRow(input string) (string, string, string) {
	parts := strings.Split(input, "=")

	start := strings.TrimSpace(parts[0])

	nodesStr := strings.ReplaceAll(parts[1], " ", "")
	nodes := strings.Split(nodesStr[1:len(nodesStr)-1], ",")

	return start, nodes[0], nodes[1]
}

func CalculateSteps(input []string) int {
	directions := input[0]

	mapping := make(map[string][2]string, 0)

	for _, line := range input[1:] {
		if line == "" {
			continue
		}

		start, left, right := parseRow(line)

		mapping[start] = [2]string{left, right}
	}

	current := "AAA"

	i := 0
	var dir string

	for current != "ZZZ" {
		dir = string(directions[i%len(directions)])

		if dir == "L" {
			current = mapping[current][0]
		} else if dir == "R" {
			current = mapping[current][1]
		} else {
			panic("Invalid direction")
		}
		i++
	}

	return i
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Println(CalculateSteps(input))
}
