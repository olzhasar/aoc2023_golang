package main

import "testing"

func TestGetFarthestPointInLoop(t *testing.T) {
	for _, test := range []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				".....",
				".S-7.",
				".|.|.",
				".L-J.",
				".....",
			},
			4,
		},
		{
			[]string{
				"..F7.",
				".FJ|.",
				"SJ.L7",
				"|F--J",
				"LJ...",
			},
			8,
		},
	} {
		actual := GetFarthestPointInLoop(test.input)

		if actual != test.expected {
			t.Errorf("Expected %d but got %d", test.expected, actual)
		}
	}
}
