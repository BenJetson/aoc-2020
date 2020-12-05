package main

import (
	"fmt"

	"github.com/BenJetson/aoc-2020/day02"
	"github.com/BenJetson/aoc-2020/utilities"
)

func main() {
	// raw stores my puzzle input.
	raw, err := utilities.ReadLinesFromFile("input.txt")
	if err != nil {
		panic(err)
	}

	passwords := make([]day02.Password, len(raw))
	for i, s := range raw {
		passwords[i] = day02.ParsePassword(s)
	}

	validCount := 0
	for _, p := range passwords {
		if p.IsSledPolicyCompliant() {
			validCount++
		}
	}

	fmt.Printf("Part one answer is: %d\n", validCount)

	validCount = 0
	for _, p := range passwords {
		if p.IsTobogganPolicyCompliant() {
			validCount++
		}
	}

	fmt.Printf("Part two answer is: %d\n", validCount)
}
