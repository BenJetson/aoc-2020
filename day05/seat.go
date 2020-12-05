package day05

import (
	"errors"
	"fmt"
)

// A Seat represents the position of a seat on the aircraft.
type Seat struct {
	// Row is the row number of the seat. Indexed at zero, values exist
	// up to RowCount (not inclusive).
	Row int
	// Col is the column number of the seat. Indexed at zero, values exist
	// up to ColCount (not inclusive).
	Col int
}

// ID returns the seat's identifier, which is equal to eight times its row
// number plus the column number.
func (s *Seat) ID() int { return s.Row*8 + s.Col }

// Binary space partition string value constants.
const (
	front = 'F'
	back  = 'B'
	left  = 'L'
	right = 'R'
)

// SeatFromBinarySpacePartition reads a string formatted as a binary space
// partition and creates the seat it represents, should it exist.
func SeatFromBinarySpacePartition(bsp string) (Seat, error) {
	if len(bsp) != 10 {
		return Seat{}, errors.New("incorrect BSP length")
	}

	// Grab the different portions of the string.
	rowPart := bsp[:7]
	colPart := bsp[7:]

	// Determine the row number using a binary search.
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

	// Determine the column number using a binary search.
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

	// These conditions ought to be mathematically impossible.
	if rowLow != rowHigh {
		return Seat{}, fmt.Errorf("row mismatch: %d != %d", rowLow, rowHigh)
	} else if colLow != colHigh {
		return Seat{}, fmt.Errorf("col mismatch: %d != %d", colLow, colHigh)
	}

	return Seat{Row: rowLow, Col: colLow}, nil
}
