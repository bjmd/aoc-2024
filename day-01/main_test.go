package main

import "testing"

func TestPartOne(t *testing.T) {
	testInput := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}

	expectedResult := 11

	result := partOne(testInput)
	if result != expectedResult {
		t.Errorf("partOne() = %d; want %d", result, expectedResult)
	}
}

func TestPartTwo(t *testing.T) {
	testInput := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}

	expectedResult := 31

	result := partTwo(testInput)
	if result != expectedResult {
		t.Errorf("partTwo() = %d; want %d", result, expectedResult)
	}
}
