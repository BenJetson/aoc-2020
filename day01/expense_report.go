package day01

import "errors"

// FindProductOfElementsWithSum takes a desired sum and a slice of integers and
// returns the product of the two elements from nums that have the desired sum.
//
// If insufficient values exist to perform this calculation or no two elements
// sum to the given value, an error is returned.
func FindProductOfElementsWithSum(sum int, nums []int) (int, error) {
	if len(nums) < 2 {
		return 0, errors.New("insufficient elements")
	}

	for outerIdx, outerVal := range nums {
		for innerIdx, innerVal := range nums {
			if outerIdx == innerIdx {
				continue
			}

			if outerVal+innerVal == sum {
				return outerVal * innerVal, nil
			}
		}
	}

	return 0, errors.New("no elements sum to total")
}

// FindProductOfThreeElementsWithSum takes a desired sum and a slice of integers
// and returns the product of the three elements from nums that have the desired
// sum.
//
// If insufficient values exist to perform this calculation or no two elements
// sum to the given value, an error is returned.
func FindProductOfThreeElementsWithSum(sum int, nums []int) (int, error) {
	if len(nums) < 3 {
		return 0, errors.New("insufficient elements")
	}

	for outerIdx, outerVal := range nums {
		for middleIdx, middleVal := range nums {
			for innerIdx, innerVal := range nums {
				if outerIdx == middleIdx ||
					middleIdx == innerIdx ||
					innerIdx == outerIdx {
					continue
				}

				if outerVal+middleVal+innerVal == sum {
					return outerVal * middleVal * innerVal, nil
				}
			}
		}
	}

	return 0, errors.New("no elements sum to total")
}
