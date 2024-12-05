package main

import "testing"

func TestPartOne(t *testing.T) {
	testInput := []string{
		"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
	}

	expectedResult := 161

	result := partOne(testInput)
	if result != expectedResult {
		t.Errorf("partOne() = %d; want %d", result, expectedResult)
	}
}

func TestPartTwo(t *testing.T) {
	testInput := []string{
		"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
	}

	expectedResult := 48

	result := partTwo(testInput)
	if result != expectedResult {
		t.Errorf("partTwo() = %d; want %d", result, expectedResult)
	}
}
