package day03

// A Mountain is two-dimensional slice of arbitrary rectangular dimensions
// that represents a grid of trees. A true position is a tree, false is blank.
type Mountain [][]bool

// IsValid determines if the mountain is truly rectangular, that is, meaning
// that all rows have the same length.
func (mtn Mountain) IsValid() bool {
	if len(mtn) < 1 {
		return false
	}

	expectLen := len(mtn[0])
	if expectLen < 1 {
		return false
	}

	for _, row := range mtn {
		if len(row) != expectLen {
			return false
		}
	}

	return true
}

// Width determines the length of a row in the mountain, given it is valid.
func (mtn Mountain) Width() int { return len(mtn[0]) }

// Height determines the number of rows in the mountain.
func (mtn Mountain) Height() int { return len(mtn) }

// ParseMountainFromLines reads a slice of lines and builds a Mountain.
// Positions containing an octothorpe are trees, all others are blank.
func ParseMountainFromLines(lines []string) Mountain {
	var mtn Mountain
	var mtnRow []bool

	for _, line := range lines {
		mtnRow = nil
		for _, char := range line {
			mtnRow = append(mtnRow, char == '#')
		}
		mtn = append(mtn, mtnRow)
	}

	return mtn
}
