package day09

import "errors"

// XMASDecoder provides useful methods for decoding "XMAS" encryption.
type XMASDecoder struct {
	Stream       []int
	PreambleSize int
}

// CheckAtPosition checks the value at the given position for validity according
// to the XMAS decoding rules.
func (dec *XMASDecoder) CheckAtPosition(pos int) (bool, error) {
	idx := pos + dec.PreambleSize
	target := dec.Stream[idx]

	if pos < 0 || idx >= len(dec.Stream) {
		return false, errors.New("out of bounds")
	}

	beginIdx := idx - dec.PreambleSize
	finalIdx := idx - 1

	pass := false
	for i := beginIdx; i <= finalIdx; i++ {
		outerVal := dec.Stream[i]
		for j := beginIdx; j <= finalIdx; j++ {
			if i == j {
				continue
			}

			innerVal := dec.Stream[j]

			if outerVal+innerVal == target {
				pass = true
			}
		}
	}

	return pass, nil
}

// FindFirstFault finds the first value in the Stream that is invalid per the
// XMAS cipher.
func (dec *XMASDecoder) FindFirstFault() (int, error) {
	for i := 0; i < len(dec.Stream)-dec.PreambleSize; i++ {
		pass, err := dec.CheckAtPosition(i)
		if err != nil {
			return -1, err
		} else if !pass {
			return dec.Stream[i+dec.PreambleSize], nil
		}
	}

	return -1, errors.New("no fault detected")
}

// FindWeaknessSumFor finds the
// FIXME
func (dec *XMASDecoder) FindWeaknessSumFor(target int) (int, error) {
	for i := dec.PreambleSize; i < len(dec.Stream)-1; i++ {
		for j := i + 1; j < len(dec.Stream); j++ {
			slice := dec.Stream[i:j]
			if sumOfSlice(slice) == target {
				minSoFar := slice[0]
				maxSoFar := slice[0]

				for _, n := range slice {
					minSoFar = min(minSoFar, n)
					maxSoFar = max(maxSoFar, n)
				}

				return minSoFar + maxSoFar, nil
			}
		}
	}
	return -1, errors.New("no weakness found")
}
