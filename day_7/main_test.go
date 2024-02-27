package main

import "testing"

func TestGetTotalWInnings(t *testing.T) {
	exampleInput := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}

	t.Run("Part 1", func(t *testing.T) {
		expected := 6440
		result := GetTotalWinnings(exampleInput)

		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}
