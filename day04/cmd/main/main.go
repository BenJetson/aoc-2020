package main

import (
	"fmt"

	"github.com/BenJetson/aoc-2020/day04"
	"github.com/BenJetson/aoc-2020/utilities"
)

func main() {
	lines, err := utilities.ReadLinesFromFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines = utilities.MergeGapsBetweenLines(lines, " ")

	var passports []day04.Passport
	for _, line := range lines {
		passports = append(passports, day04.ParsePassportLine(line))
	}

	count := 0
	extCount := 0
	for _, p := range passports {
		if p.HasRequiredFields() {
			count++
		}

		if p.IsValid() {
			extCount++
		}
	}

	fmt.Printf("Part one answer is: %d\n", count)
	fmt.Printf("Part two answer is: %d\n", extCount)
}
