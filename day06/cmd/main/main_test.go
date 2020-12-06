package main

import (
	"testing"

	"github.com/BenJetson/aoc-2020/utilities"
)

func TestMain(t *testing.T) {
	// The solutions to both parts of the puzzle.
	part1 := 6596
	part2 := 3219

	// Run the driver and verify the console output.
	utilities.CheckMainPuzzleSolution(t, part1, part2)
}
