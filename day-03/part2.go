package main

import (
	"regexp"
	"strconv"
)

func partTwo(lines []string) int {

	var returnVal int
	mulPattern := `^mul\((\d{1,3}),(\d{1,3})\)`
	dontPattern := `^don\'t\(\)`
	doPattern := `^do\(\)`

	for _, line := range lines {
		var dont bool
		var do bool

		do = true

		for len(line) > 0 {
			// Regex search the start f each line for the matching pattern
			r := regexp.MustCompile(mulPattern)
			mulMatch := r.FindAllStringSubmatch(line, -1)

			r = regexp.MustCompile(dontPattern)
			dontMatch := r.FindAllStringSubmatch(line, -1)

			r = regexp.MustCompile(doPattern)
			doMatch := r.FindAllStringSubmatch(line, -1)

			// If the multiply matches then add it and remove from the line.
			if len(mulMatch) > 0 && do && !dont {
				mulLength := len(mulMatch[0][0])
				num1, _ := strconv.Atoi(mulMatch[0][1])
				num2, _ := strconv.Atoi(mulMatch[0][2])
				returnVal += num1 * num2
				line = line[mulLength:]
				continue
			}

			// if dont matches then update bools and remove from the line
			if len(dontMatch) > 0 {
				do = false
				dont = true
				line = line[7:]
				continue
			}

			// if do matches then update bools and remove from the line
			if len(doMatch) > 0 {
				do = true
				dont = false
				line = line[4:]
				continue
			}

			// If nothing matches then simply step forward one character
			line = line[1:]
		}
	}

	return returnVal
}
