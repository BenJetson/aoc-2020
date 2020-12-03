package day03

type Mountain [][]bool

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

func (mtn Mountain) Width() int  { return len(mtn[0]) }
func (mtn Mountain) Height() int { return len(mtn) }

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
