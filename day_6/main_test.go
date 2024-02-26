package main

import "testing"

func TestGetNumWaysToWinMultiplied(t *testing.T) {
	exampleInput := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}

	expected := 288
	result := GetNumWaysToWinMultiplied(exampleInput)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
