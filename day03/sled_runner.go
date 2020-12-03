package day03

type SledRunner struct {
	mtn        Mountain
	slope      Slope
	mtnW, mtnH int
	posX, posY int
}

func MakeSledRunner(mtn Mountain, slope Slope) SledRunner {
	return SledRunner{
		mtn:   mtn,
		slope: slope,
		mtnW:  mtn.Width(),
		mtnH:  mtn.Height(),
	}
}

func (s *SledRunner) Advance() bool {
	s.posX = (s.posX + s.slope.DeltaX) % s.mtnW
	s.posY = (s.posY + s.slope.DeltaY)

	return s.posY < s.mtnH
}

func (s *SledRunner) Reset() {
	s.posX = 0
	s.posY = 0
}

func (s *SledRunner) HitsTree() bool {
	return s.mtn[s.posY][s.posX]
}

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
