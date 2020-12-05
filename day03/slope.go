package day03

// A Slope represents the amount of change a skiier shall take per timeframe.
// It is important to note the modified coordinate plane directionality due
// to row-major order.
type Slope struct {
	// DeltaX is the amount the X coordinate should change.
	// Negative is left, positive is right.
	DeltaX int
	// DeltaY is the amount the Y coordinate should change.
	// Positive is down, negative is up.
	DeltaY int
}
