package day07

import (
	"strconv"
	"strings"
)

// BagColor is an alias of string, used as a safety to ensure that colors are
// in the right spot.
type BagColor string

// A BagRuleIndex is an map from a color of bag to the rule detailing its
// required contents.
type BagRuleIndex map[BagColor]BagRule

// A BagRule specifies how many bags of a given color must be inside this bag.
type BagRule map[BagColor]int

// CanDirectlyContainColor tells whether or not this bag may directly contain
// the indicated target color.
func (rule BagRule) CanDirectlyContainColor(target BagColor) bool {
	_, ok := rule[target]
	return ok
}

// ParseBagRuleIndex takes a slice of bag rule lines and transforms them into a
// bag rule index.
func ParseBagRuleIndex(lines []string) (BagRuleIndex, error) {
	ruleIndex := make(BagRuleIndex)
	for _, line := range lines {
		rule := make(BagRule)

		breakIndex := strings.Index(line, " bags contain ")
		outer := BagColor(line[:breakIndex])

		line = line[breakIndex+14 : len(line)-1]

		if line == "no other bags" {
			ruleIndex[outer] = rule
			continue
		}

		contentLines := strings.Split(line, ", ")

		for _, cLine := range contentLines {
			breakIndex = strings.Index(cLine, " ")
			count, err := strconv.Atoi(cLine[:breakIndex])
			if err != nil {
				return nil, err
			}

			lastIndex := strings.Index(cLine, " bag")
			inner := BagColor(cLine[breakIndex+1 : lastIndex])

			rule[inner] = count
		}

		ruleIndex[outer] = rule
	}

	return ruleIndex, nil
}

// CountCanContainAtLeastOneOf determines the number of bag colors that may
// contain at least one of the target color. If a given bag type may not contain
// the target color directly, a recursive methodology is used to determine if
// any bags it contains may contain the target color.
func (ruleIndex BagRuleIndex) CountCanContainAtLeastOneOf(target BagColor) int {
	var count int
	for current, rule := range ruleIndex {
		if rule.CanDirectlyContainColor(target) || ruleIndex.recursiveCanContain(current, target) {
			count++
			continue
		}
	}

	return count
}

func (ruleIndex BagRuleIndex) recursiveCanContain(current, target BagColor) bool {
	currentRule := ruleIndex[current]
	if len(currentRule) == 0 {
		return false
	} else if currentRule.CanDirectlyContainColor(target) {
		return true
	}

	for color := range currentRule {
		if ruleIndex.recursiveCanContain(color, target) {
			return true
		}
	}

	return false
}

// TargetMustContainExactly uses a recursive methodology to determine the
// number of bags that the target color of bag must contain to pass the rule.
func (ruleIndex BagRuleIndex) TargetMustContainExactly(target BagColor) int {
	rule := ruleIndex[target]

	count := 0

	for color, innerCount := range rule {
		count += innerCount
		count += innerCount * ruleIndex.TargetMustContainExactly(color)
	}

	return count
}
