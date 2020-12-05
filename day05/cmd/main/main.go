package main

import (
	"fmt"

	"github.com/BenJetson/aoc-2020/day05"
	"github.com/BenJetson/aoc-2020/utilities"
)

func main() {
	// My puzzle input.
	lines, err := utilities.ReadLinesFromFile("input.txt")
	if err != nil {
		panic(err)
	}

	var s day05.Seat
	var seats []day05.Seat
	for _, line := range lines {
		s, err = day05.SeatFromBinarySpacePartition(line)
		if err != nil {
			panic(err)
		}

		seats = append(seats, s)
	}

	var maxSoFar = seats[0].ID()
	for _, seat := range seats {
		id := seat.ID()
		if id > maxSoFar {
			maxSoFar = id
		}
	}

	fmt.Printf("Part one answer is: %d\n", maxSoFar)

	m := new(day05.SeatMap)
	for _, seat := range seats {
		m.MarkSeatTaken(seat)
	}

	mine, err := m.FindEmptySeatInPartialRow()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part two answer is: %d\n", mine.ID())
}
