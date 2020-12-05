package day05

import (
	"errors"
	"fmt"
)

type Seat struct {
	Row int
	Col int
}

func (s *Seat) ID() int { return s.Row*8 + s.Col }

const front = 'F'
const back = 'B'
const left = 'L'
const right = 'R'

func SeatFromBinarySpacePartition(bsp string) (Seat, error) {
	if len(bsp) != 10 {
		return Seat{}, errors.New("incorrect BSP length")
	}

	rowPart := bsp[:7]
	colPart := bsp[7:]

	rowLow := 0
	rowHigh := RowCount - 1

	for _, c := range rowPart {
		diff := (rowHigh - rowLow) / 2

		if c == front {
			rowHigh = rowLow + diff
		} else if c == back {
			rowLow = rowHigh - diff
		} else {
			return Seat{}, fmt.Errorf("invalid row character '%v'", string(c))
		}
	}

	colLow := 0
	colHigh := ColCount - 1

	for _, c := range colPart {
		diff := (colHigh - colLow) / 2

		if c == left {
			colHigh = colLow + diff
		} else if c == right {
			colLow = colHigh - diff
		} else {
			return Seat{}, fmt.Errorf("invalid col character '%v'", string(c))
		}
	}

	if rowLow != rowHigh {
		return Seat{}, fmt.Errorf("row mismatch: %d != %d", rowLow, rowHigh)
	} else if colLow != colHigh {
		return Seat{}, fmt.Errorf("col mismatch: %d != %d", rowLow, rowHigh)
	}

	return Seat{Row: rowLow, Col: colLow}, nil
}
