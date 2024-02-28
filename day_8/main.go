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

func makeMap(input []string) map[string][2]string {
	mapping := make(map[string][2]string, 0)

	for _, line := range input {
		if line == "" {
			continue
		}

		start, left, right := parseRow(line)

		mapping[start] = [2]string{left, right}
	}

	return mapping
}

func CalculateSteps(input []string) int {
	directions := input[0]

	mapping := makeMap(input[1:])

	current := "AAA"

	i := 0
	var dir byte

	for current != "ZZZ" {
		dir = directions[i%len(directions)]

		if dir == 'L' {
			current = mapping[current][0]
		} else if dir == 'R' {
			current = mapping[current][1]
		} else {
			panic("Invalid direction")
		}
		i++
	}

	return i
}

func findPathLength(directions string, mapping map[string][2]string, start string) int {
	i := 0
	var dir byte

	for start[2] != 'Z' {
		dir = directions[i%len(directions)]

		if dir == 'L' {
			start = mapping[start][0]
		} else if dir == 'R' {
			start = mapping[start][1]
		} else {
			panic("Invalid direction")
		}
		i++
	}

	return i
}

func calculateGCD(a, b uint64) uint64 {
	// Euclidean algorithm
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func calculateLCM(a, b uint64) uint64 {
	gcd := calculateGCD(a, b)
	return a * b / gcd
}

func CalculateStepsPartTwo(input []string) uint64 {
	directions := input[0]

	mapping := makeMap(input[1:])

	var paths []int

	for k := range mapping {
		if k[2] == 'A' {
			paths = append(paths, findPathLength(directions, mapping, k))
		}
	}

	lcm := uint64(paths[0])

	for _, path := range paths[1:] {
		lcm = calculateLCM(lcm, uint64(path))
	}

	return lcm
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

	fmt.Println(CalculateStepsPartTwo(input))
}
