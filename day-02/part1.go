package main

import (
	"math"
	"strconv"
	"strings"
)

func diff(a, b float64) float64 {
	if a < b {
		return b - a
	}
	return a - b
}

// Return true for decrease and false for increase
func checkIncrement(list []float64) bool {
	var prev bool

	for i := range len(list) {
		if i+1 == len(list) {
			break
		}
		diff := diff(list[i], list[i+1])
		if diff < 1 || diff > 3 {
			return false
		}

		sign := math.Signbit(list[i] - list[i+1])
		if i == 0 {
			prev = sign
			continue
		}
		if prev != sign {
			return false
		}

	}
	return true
}

func partOne(lines []string) int {

	var safeCount int

	for _, line := range lines {
		cols := strings.Fields(line)
		reports := make([]float64, 0, len(lines))

		for _, a := range cols {
			a, _ := strconv.ParseFloat(a, 64)
			reports = append(reports, a)
		}

		safe := checkIncrement(reports)

		if safe {
			safeCount += 1
		}
	}

	return safeCount
}
