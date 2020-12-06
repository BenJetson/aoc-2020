package main

import (
	"testing"

	"github.com/BenJetson/aoc-2020/utilities"
)

func TestMain(t *testing.T) {
	// The solutions to both parts of the puzzle.
	part1 := 262
	part2 := 2698900776

	// Run the driver and verify the console output.
	utilities.CheckMainPuzzleSolution(t, part1, part2)
}
