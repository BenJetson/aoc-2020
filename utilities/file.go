package utilities

import (
	"io/ioutil"
)

const linefeed = '\n'

// ReadLinesFromFile reads the entire file specified, and returns a slice
// of strings
func ReadLinesFromFile(filename string) ([]string, error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var lines []string
	var line string

	for _, b := range raw {
		if b == linefeed {
			lines = append(lines, line)
			line = ""

			continue
		}

		line += string(b)
	}

	return lines, nil
}

// ReadIntegersFromFile takes a
func ReadIntegersFromFile(filename string) ([]int, error) {
	lines, err := ReadLinesFromFile(filename)
	if err != nil {
		return nil, err
	}

	nums, err := SliceStringsToInts(lines)
	if err != nil {
		return nil, err
	}

	return nums, nil
}
