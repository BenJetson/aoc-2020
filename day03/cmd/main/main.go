package main

import (
	"fmt"

	"github.com/BenJetson/aoc-2020/day03"
	"github.com/BenJetson/aoc-2020/utilities"
)

func main() {
	lines, err := utilities.ReadLinesFromFile("input.txt")
	if err != nil {
		panic(err)
	}

	mtn := day03.ParseMountainFromLines(lines)
	slope := day03.Slope{DeltaX: 3, DeltaY: 1}
	sr := day03.MakeSledRunner(mtn, slope)

	count := sr.CountTreesOnRoute()

	fmt.Printf("Part one answer is: %d\n", count)

	total := 1
	slopes := []day03.Slope{
		{DeltaX: 1, DeltaY: 1},
		{DeltaX: 3, DeltaY: 1},
		{DeltaX: 5, DeltaY: 1},
		{DeltaX: 7, DeltaY: 1},
		{DeltaX: 1, DeltaY: 2},
	}

	for _, slope = range slopes {
		sr = day03.MakeSledRunner(mtn, slope)
		total *= sr.CountTreesOnRoute()
	}

	fmt.Printf("Part two answer is: %d\n", total)
}
