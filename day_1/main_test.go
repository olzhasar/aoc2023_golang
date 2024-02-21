package main

import (
	"testing"
)

func TestCalculateCalibration(t *testing.T) {
	exampleInput := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	expected := []int{12, 38, 15, 77}

	for i, str := range exampleInput {
		result := calculateCalibration(str)
		if result != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], result)
		}
	}
}

func TestCalculateCalibrationPartTwo(t *testing.T) {
	exampleInput := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	expected := []int{29, 83, 13, 24, 42, 14, 76}

	for i, str := range exampleInput {
		result := calculateCalibrationPartTwo(str)
		if result != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], result)
		}
	}
}
