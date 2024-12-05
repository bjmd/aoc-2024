package main

import (
	"regexp"
	"strconv"
	"strings"
)

func partTwo(lines []string) int {

	var returnVal int

	// Join all lines into a single string
	input := strings.Join(lines, "")
	enabled := true

	// Case/Switch implementation. Cleaner and more concise but around the same execution time.
	pattern := `(mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\))`
	r := regexp.MustCompile(pattern)
	matches := r.FindAllStringSubmatch(input, -1)

	// Process matches in order
	for _, match := range matches {
		switch match[1] {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				num1, _ := strconv.Atoi(match[2])
				num2, _ := strconv.Atoi(match[3])
				returnVal += num1 * num2
			}
		}
	}

	// Original implementation
	//
	// Find all matches at once
	// mulR := regexp.MustCompile(`^mul\((\d{1,3}),(\d{1,3})\)`)
	// dontR := regexp.MustCompile(`^don\'t\(\)`)
	// doR := regexp.MustCompile(`^do\(\)`)

	// for len(input) > 0 {
	// 	if match := mulR.FindStringSubmatch(input); match != nil && enabled {
	// 		num1, _ := strconv.Atoi(match[1])
	// 		num2, _ := strconv.Atoi(match[2])
	// 		returnVal += num1 * num2
	// 		input = input[len(match[0]):]
	// 		continue
	// 	}

	// 	if match := dontR.FindStringSubmatch(input); match != nil {
	// 		enabled = false
	// 		input = input[7:]
	// 		continue
	// 	}

	// 	if match := doR.FindStringSubmatch(input); match != nil {
	// 		enabled = true
	// 		input = input[4:]
	// 		continue
	// 	}

	// 	input = input[1:]
	// }

	return returnVal
}
