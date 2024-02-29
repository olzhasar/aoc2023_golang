package main

import "testing"

func TestGetExtrapolatedSum(t *testing.T) {
	input := []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}
	expected := 114

	actual := GetExtrapolatedSum(input)

	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
