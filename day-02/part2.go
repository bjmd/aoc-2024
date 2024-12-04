package main

import (
	"math"
	"strconv"
	"strings"
)

func diff2(a, b float64) float64 {
	if a < b {
		return b - a
	}
	return a - b
}

// Return true for decrease and false for increase
func checkIncrement2(list []float64) (bool, int) {
	// this bool catches the first step and whether it is an increase or decrease.
	// This is then used for the comparison for the remaining steps.
	var increment bool

	for i := range len(list) {
		if i+1 == len(list) {
			break
		}

		diff := diff(list[i], list[i+1])
		if diff < 1 || diff > 3 {
			return false, i
		}

		sign := math.Signbit(list[i] - list[i+1])
		if i == 0 {
			increment = sign
			continue
		}
		if increment != sign {
			return false, i
		}

	}
	return true, 0
}

func partTwo(lines []string) int {

	var safeCount int
	var prevFail bool

	for _, line := range lines {
		cols := strings.Fields(line)
		report := make([]float64, 0, len(lines))

		for _, a := range cols {
			a, _ := strconv.ParseFloat(a, 64)
			report = append(report, a)
		}

		origReport := make([]float64, len(report))
		copy(origReport, report)
		var retryIdx []int

		for {
			safe, idx := checkIncrement2(report)
			if safe {
				prevFail = false
				safeCount += 1
				break
			}
			if len(retryIdx) == 0 {
				if prevFail {
					prevFail = false
					break
				}

				retryIdx = append(retryIdx, idx+1)

				retryIdx = append(retryIdx, idx)

				if idx != 0 {
					retryIdx = append(retryIdx, idx-1)
				}

			}

			report = make([]float64, len(origReport))
			copy(report, origReport)
			report = append(report[:retryIdx[0]], report[retryIdx[0]+1:]...)
			retryIdx = retryIdx[1:]
			prevFail = true
		}
	}

	return safeCount
}
