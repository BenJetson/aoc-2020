package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSeatFromBinarySpacePartition(t *testing.T) {
	testCases := []struct {
		alias      string
		bsp        string
		expectSeat Seat
		expectID   int
		expectErr  bool
	}{
		{
			alias: "given1",
			bsp:   "BFFFBBFRRR",
			expectSeat: Seat{
				Row: 70,
				Col: 7,
			},
			expectID: 567,
		},
		{
			alias: "given2",
			bsp:   "FFFBBBFRRR",
			expectSeat: Seat{
				Row: 14,
				Col: 7,
			},
			expectID: 119,
		},
		{
			alias: "given3",
			bsp:   "BBFFBBFRLL",
			expectSeat: Seat{
				Row: 102,
				Col: 4,
			},
			expectID: 820,
		},
		{
			alias:     "badRow",
			bsp:       "BBFFBXFRLL",
			expectErr: true,
		},
		{
			alias:     "badCol",
			bsp:       "BBFFBBFRXL",
			expectErr: true,
		},
		{
			alias:     "badLen",
			bsp:       "BBLLR",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.alias, func(t *testing.T) {
			seat, err := SeatFromBinarySpacePartition(tc.bsp)

			if tc.expectErr {
				require.Error(t, err)
				assert.Zero(t, seat)
			} else {
				require.NoError(t, err)
				assert.EqualValues(t, tc.expectSeat, seat)
				assert.Equal(t, tc.expectID, seat.ID())
			}
		})
	}
}
