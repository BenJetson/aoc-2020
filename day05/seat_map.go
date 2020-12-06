package day05

import (
	"errors"
)

const (
	// RowCount is the maximum number of rows on an aircraft.
	RowCount = 128
	// ColCount is the maximum number of columns on an aircraft.
	ColCount = 8
)

// A SeatMap represents the taken/available status of every single seat aboard
// the aircraft.
type SeatMap [RowCount][ColCount]bool

// MarkSeatTaken marks the given seat as taken.
func (m *SeatMap) MarkSeatTaken(s Seat) {
	m[s.Row][s.Col] = true
}

// FindEmptySeatInPartialRow searches the map for a row of seats that is both
// not empty and not full. The first such available seat from the front of the
// plane, left to right, is returned if such exists.
func (m *SeatMap) FindEmptySeatInPartialRow() (Seat, error) {
	for i := 0; i < RowCount; i++ {
		// This aircraft is missing some rows towards the front and back, and
		// we are guaranteed to be the only empty seat. Therefore, if a given
		// side of the aisle is completely empty we know it does not exist.
		var isLeftPartial, isRightPartial bool
		for j, taken := range m[i] {
			isLeftPartial = isLeftPartial || (taken && j < ColCount/2)
			isRightPartial = isRightPartial || (taken && j >= ColCount/2)
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
