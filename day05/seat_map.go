package day05

import (
	"errors"
)

const RowCount = 128
const ColCount = 8

type SeatMap [RowCount][ColCount]bool

func (m *SeatMap) MarkSeatTaken(s Seat) {
	m[s.Row][s.Col] = true
}

func (m *SeatMap) FindEmptySeatInPartialRow() (Seat, error) {
	for i := 0; i < RowCount; i++ {
		// This aircraft is missing some rows towards the front and back, and
		// we are guaranteed to be the only empty seat. Therefore, if a given
		// side of the aisle is completely empty we know it does not exist.
		isLeftPartial := false
		isRightPartial := false
		for j, taken := range m[i] {
			if j < ColCount/2 {
				if taken {
					isLeftPartial = true
				}
			} else {
				if taken {
					isRightPartial = true
				}
			}
		}

		if isLeftPartial {
			// Check for empty seats on left side of aisle.
			for j := 0; j < ColCount/2; j++ {
				if !m[i][j] {
					return Seat{
						Row: i,
						Col: j,
					}, nil
				}
			}
		}

		if isRightPartial {
			// Check for empty seats on right side of aisle.
			for j := ColCount / 2; j < ColCount; j++ {
				if !m[i][j] {
					return Seat{
						Row: i,
						Col: j,
					}, nil
				}
			}
		}
	}

	return Seat{}, errors.New("no empty seats available in partial rows")
}
