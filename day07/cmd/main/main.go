package main

import (
	"fmt"

	"github.com/BenJetson/aoc-2020/day07"
	"github.com/BenJetson/aoc-2020/utilities"
)

func main() {
	// Read puzzle input from file.
	lines, err := utilities.ReadLinesFromFile("input.txt")
	if err != nil {
		panic(err)
	}

	// Transform input into an index of rules.
	ruleIndex, err := day07.ParseBagRuleIndex(lines)
	if err != nil {
		panic(err)
	}

	count := ruleIndex.CountCanContainAtLeastOneOf("shiny gold")
	fmt.Printf("Part one answer is: %d\n", count)

	count = ruleIndex.TargetMustContainExactly("shiny gold")
	fmt.Printf("Part two answer is: %d\n", count)
}
