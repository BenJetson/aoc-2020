package main

import (
	"fmt"

	"github.com/BenJetson/aoc-2020/day08"
	"github.com/BenJetson/aoc-2020/utilities"
)

func main() {
	lines, err := utilities.ReadLinesFromFile("input.txt")
	if err != nil {
		panic(err)
	}

	p, err := day08.ParseProgram(lines)
	if err != nil {
		panic(err)
	}

	acc, err := p.RunUntilLoop()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part one answer is: %d\n", acc)

	acc, err = p.FixOneBugInfiniteLoop()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part two answer is: %d\n", acc)
}
