package main

import "testing"

func TestPartOne(t *testing.T) {
	testInput := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}

	expectedResult := 2

	result := partOne(testInput)
	if result != expectedResult {
		t.Errorf("partOne() = %d; want %d", result, expectedResult)
	}
}

func TestPartTwo(t *testing.T) {
	testInput := []string{
		// "7 6 4 2 1",
		// "1 2 7 8 9",
		// "9 7 6 2 1",
		// "1 3 2 4 5",
		// "8 6 4 4 1",
		// "1 3 6 7 9",
		"23 26 23 24 27 28",
	}

	expectedResult := 4

	result := partTwo(testInput)
	if result != expectedResult {
		t.Errorf("partTwo() = %d; want %d", result, expectedResult)
	}
}
