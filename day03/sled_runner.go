package day03

// A SledRunner provides utilities for navigating a sled across a mountain
// along a fixed linear path.
type SledRunner struct {
	mtn        Mountain
	slope      Slope
	mtnW, mtnH int
	posX, posY int
}

// MakeSledRunner creates a new SledRunner given the mountain and slope.
func MakeSledRunner(mtn Mountain, slope Slope) SledRunner {
	return SledRunner{
		mtn:   mtn,
		slope: slope,
		mtnW:  mtn.Width(),
		mtnH:  mtn.Height(),
	}
}

// Advance moves the sled one slope unit. Once the bottom of the mountain
// has been reached, this method returns false and the sled is left in an
// invalid position.
func (s *SledRunner) Advance() bool {
	s.posX = (s.posX + s.slope.DeltaX) % s.mtnW
	s.posY = (s.posY + s.slope.DeltaY)

	return s.posY < s.mtnH
}

// Reset will reset the position of the sled to the top left corner (0, 0).
func (s *SledRunner) Reset() {
	s.posX = 0
	s.posY = 0
}

// HitsTree determines whether or not the sled has hit a tree at its current
// position on the mountain.
func (s *SledRunner) HitsTree() bool {
	return s.mtn[s.posY][s.posX]
}

// CountTreesOnRoute will advance the sled until the sled reaches the bottom
// of hte mountain, counting the number of trees hit along the route.
func (s *SledRunner) CountTreesOnRoute() int {
	s.Reset()

	count := 0
	for s.Advance() {
		if s.HitsTree() {
			count++
		}
	}

	return count
}
