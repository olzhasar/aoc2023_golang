package main

import "testing"

func TestGetNumWaysToWinMultiplied(t *testing.T) {
	exampleInput := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}

	t.Run("Part 1", func(t *testing.T) {
		expected := 288
		result := GetNumWaysToWinMultiplied(exampleInput)

		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("Part 2", func(t *testing.T) {
		expected := 71503
		result := GetNumWaysToWinMultipliedPartTwo(exampleInput)

		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}
