package main

import (
	"fmt"

	"github.com/BenJetson/aoc-2020/day01"
	"github.com/BenJetson/aoc-2020/utilities"
)

func main() {
	// nums holds my puzzle input.
	nums, err := utilities.ReadIntegersFromFile("input.txt")
	if err != nil {
		panic(err)
	}

	sum := 2020

	answer, err := day01.FindProductOfElementsWithSum(sum, nums)

	if err != nil {
		panic("No solution for part one!")
	}

	fmt.Printf("Part one answer is: %d\n", answer)

	answer, err = day01.FindProductOfThreeElementsWithSum(sum, nums)

	if err != nil {
		panic("No solution for part two!")
	}

	fmt.Printf("Part two answer is: %d\n", answer)
}
