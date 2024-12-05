package main

import (
	"regexp"
	"strconv"
)

func partOne(lines []string) int {

	var returnVal int
	mulPattern := `^mul\((\d{1,3}),(\d{1,3})\)`

	for _, line := range lines {

		for len(line) > 0 {
			r := regexp.MustCompile(mulPattern)
			matches := r.FindAllStringSubmatch(line, -1)

			if len(matches) > 0 {
				mulLength := len(matches[0][0])
				num1, _ := strconv.Atoi(matches[0][1])
				num2, _ := strconv.Atoi(matches[0][2])
				returnVal += num1 * num2
				line = line[mulLength:]
				continue
			}
			line = line[1:]

		}
	}

	return returnVal
}
