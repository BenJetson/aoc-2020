package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindProductOfElementsWithSum(t *testing.T) {
	testCases := []struct {
		alias     string
		sum       int
		nums      []int
		expectVal int
		expectErr bool
	}{
		{
			alias:     "nil_values",
			sum:       299,
			nums:      nil,
			expectErr: true,
		},
		{
			alias:     "empty_values",
			sum:       1023,
			nums:      []int{},
			expectErr: true,
		},
		{
			alias:     "insufficient_values",
			sum:       3920,
			nums:      []int{29},
			expectErr: true,
		},
		{
			alias:     "no_sum",
			sum:       881,
			nums:      []int{29, 6, 7, 810, 392, 684, 444, 0, 0, 1020, 5, 100},
			expectErr: true,
		},
		{
			alias:     "sample",
			sum:       2020,
			nums:      []int{1721, 979, 366, 299, 675, 1456},
			expectVal: 514579,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.alias, func(t *testing.T) {
			val, err := FindProductOfElementsWithSum(tc.sum, tc.nums)

			if !tc.expectErr {
				require.NoError(t, err)
				assert.Equal(t, tc.expectVal, val)
			} else {
				require.Error(t, err)
			}
		})
	}
}
func TestFindProductOfThreeElementsWithSum(t *testing.T) {
	testCases := []struct {
		alias     string
		sum       int
		nums      []int
		expectVal int
		expectErr bool
	}{
		{
			alias:     "nil_values",
			sum:       299,
			nums:      nil,
			expectErr: true,
		},
		{
			alias:     "empty_values",
			sum:       1023,
			nums:      []int{},
			expectErr: true,
		},
		{
			alias:     "insufficient_values",
			sum:       3920,
			nums:      []int{29, 4},
			expectErr: true,
		},
		{
			alias:     "no_sum",
			sum:       881,
			nums:      []int{29, 6, 7, 810, 392, 684, 444, 0, 0, 1020, 5, 100},
			expectErr: true,
		},
		{
			alias:     "sample",
			sum:       2020,
			nums:      []int{1721, 979, 366, 299, 675, 1456},
			expectVal: 241861950,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.alias, func(t *testing.T) {
			val, err := FindProductOfThreeElementsWithSum(tc.sum, tc.nums)

			if !tc.expectErr {
				require.NoError(t, err)
				assert.Equal(t, tc.expectVal, val)
			} else {
				require.Error(t, err)
			}
		})
	}
}
