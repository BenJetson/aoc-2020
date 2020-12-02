package day02

import (
	"fmt"
	"strings"
)

// Password stores metadata about one user's password from a Toboggan company
// database file.
type Password struct {
	PosA   int
	PosB   int
	Target string
	Value  string
}

// ParsePassword creates a new Password struct given a line from the Toboggan
// company database.
func ParsePassword(s string) Password {
	var p Password
	fmt.Sscanf(s, "%d-%d %1s: %s", &p.PosA, &p.PosB, &p.Target, &p.Value)
	return p
}

// IsSledPolicyCompliant determines if a password is compliant with the sled
// company password policy.
func (p *Password) IsSledPolicyCompliant() bool {
	count := strings.Count(p.Value, p.Target)
	return count >= p.PosA && count <= p.PosB
}

// IsTobogganPolicyCompliant determines if a password is compliant with the
// Toboggan company password policy.
func (p *Password) IsTobogganPolicyCompliant() bool {
	// Toboggan policy indexes at 1, but Go strings index (sanely) at 0.
	idxA := p.PosA - 1
	idxB := p.PosB - 1

	// Check bounds on password value before subscripting.
	l := len(p.Value) - 1
	if l < idxA || l < idxB {
		return false
	}

	// Determine if values at positions A and B match the target rune.
	target := rune(p.Target[0])
	matchesA := []rune(p.Value)[idxA] == target
	matchesB := []rune(p.Value)[idxB] == target

	// Is compliant if (matchesA XOR matchesB)
	return (matchesA || matchesB) && !(matchesA && matchesB)
}
