package main

import "testing"

func TestGetPartNumbersSum(t *testing.T) {
	exampleInput := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	expected := 4361

	result := getPartNumbersSum(exampleInput)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestGetGearRatioSum(t *testing.T) {
	exampleInput := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	expected := 467835

	result := getGearRatioSum(exampleInput)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
