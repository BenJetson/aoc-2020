package main

import (
	"fmt"

	"github.com/BenJetson/aoc-2020/day06"
	"github.com/BenJetson/aoc-2020/utilities"
)

func main() {
	// My puzzle input.
	lines, err := utilities.ReadLinesFromFile("input.txt")
	if err != nil {
		panic(err)
	}

	surveyGroups := day06.ParseLinesToSurveyGroups(lines)

	var answeredSum, commonSum int
	for _, sg := range surveyGroups {
		answeredSum += sg.FindAnsweredQuestionCount()
		commonSum += sg.FindCommonQuestionCount()
	}

	fmt.Printf("Part one answer is: %d\n", answeredSum)
	fmt.Printf("Part two answer is: %d\n", commonSum)
}
