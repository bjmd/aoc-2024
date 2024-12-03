package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func partOne(lines []string) {
	var col1 []int
	var col2 []int
	var diff int

	for _, line := range lines {
		cols := strings.Fields(line)
		num1, _ := strconv.Atoi(cols[0])
		num2, _ := strconv.Atoi(cols[1])

		col1 = append(col1, num1)
		col2 = append(col2, num2)
	}

	sort.Ints(col1)
	sort.Ints(col2)

	for i := range len(lines) {

		if col1[i] > col2[i] {
			diff += col1[i] - col2[i]
		}

		if col1[i] < col2[i] {
			diff += col2[i] - col1[i]
		}

	}

	fmt.Println("Total diff: ", diff)
}
