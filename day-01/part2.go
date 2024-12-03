package main

import (
	"strconv"
	"strings"
)

func partTwo(lines []string) int {
	var col1 []int
	var col2 []int
	var total int

	for _, line := range lines {
		cols := strings.Fields(line)
		num1, _ := strconv.Atoi(cols[0])
		num2, _ := strconv.Atoi(cols[1])

		col1 = append(col1, num1)
		col2 = append(col2, num2)
	}

	for _, i := range col1 {
		var occurrences int
		for _, x := range col2 {
			if i == x {
				occurrences += 1
			}
		}

		total += i * occurrences
	}

	return total
}
