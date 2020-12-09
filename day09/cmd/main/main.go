package main

import (
	"fmt"

	"github.com/BenJetson/aoc-2020/day09"
	"github.com/BenJetson/aoc-2020/utilities"
)

func main() {
	stream, err := utilities.ReadIntegersFromFile("input.txt")
	if err != nil {
		panic(err)
	}

	dec := day09.XMASDecoder{
		Stream:       stream,
		PreambleSize: 25,
	}

	val, err := dec.FindFirstFault()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part one answer is: %d\n", val)

	sum, err := dec.FindWeaknessSumFor(val)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part two answer is: %d\n", sum)
}
