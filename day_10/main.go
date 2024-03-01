package main

import (
	"bufio"
	"fmt"
	"os"
)

type ORIGIN int

const (
	START ORIGIN = iota
	NORTH
	SOUTH
	WEST
	EAST
)

func makeMatrix(input []string) [][]rune {
	var matrix [][]rune
	for _, line := range input {
		matrix = append(matrix, []rune(line))
	}
	return matrix
}

func getStart(matrix [][]rune) (int, int) {
	for y, row := range matrix {
		for x, val := range row {
			if val == 'S' {
				return x, y
			}
		}
	}
	panic("No start found")
}

type Cell struct {
	y, x   int
	origin ORIGIN
}

func isLoopConnected(matrix [][]rune, cell Cell) bool {
	if cell.origin == START {
		return true
	}

	r := matrix[cell.y][cell.x]

	if r == 'S' {
		return true
	}

	switch cell.origin {
	case NORTH:
		return r == '|' || r == 'J' || r == 'L'
	case SOUTH:
		return r == '|' || r == '7' || r == 'F'
	case WEST:
		return r == '-' || r == 'J' || r == '7'
	case EAST:
		return r == '-' || r == 'L' || r == 'F'
	}

	panic("Invalid cell")
}

func getAdjacentCells(matrix [][]rune, x, y int) []Cell {
	var cells []Cell

	appendIfValid := func(newX, newY int, origin ORIGIN) {
		cell := Cell{newY, newX, origin}
		if isLoopConnected(matrix, cell) {
			cells = append(cells, cell)
		}
	}

	if x > 0 {
		appendIfValid(x-1, y, EAST)
	}

	if x < len(matrix[0])-1 {
		appendIfValid(x+1, y, WEST)
	}

	if y > 0 {
		appendIfValid(x, y-1, SOUTH)
	}

	if y < len(matrix)-1 {
		appendIfValid(x, y+1, NORTH)
	}

	return cells
}

func getMaxPathLength(matrix [][]rune, visited [][]bool, cell Cell, current int) int {
	r := matrix[cell.y][cell.x]

	if cell.origin != START && r == 'S' {
		return current
	}

	if r != 'S' {
		visited[cell.y][cell.x] = true
	}

	maxLength := 0

	for _, c := range getAdjacentCells(matrix, cell.x, cell.y) {
		if visited[c.y][c.x] == true {
			continue
		}
		maxLength = max(maxLength, getMaxPathLength(matrix, visited, c, current+1))
	}

	visited[cell.y][cell.x] = false

	return maxLength
}

func GetFarthestPointInLoop(input []string) int {
	matrix := makeMatrix(input)

	visited := make([][]bool, len(matrix))

	for i := range matrix {
		visited[i] = make([]bool, len(matrix[i]))
	}

	startX, startY := getStart(matrix)

	maxPathLength := getMaxPathLength(matrix, visited, Cell{startY, startX, START}, 0)

	return maxPathLength / 2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

	fmt.Println(GetFarthestPointInLoop(input))
}
