package ent

type horizontal struct {
	line
}

func (s *horizontal) makeLine(jumble *[][]IdxPosition) *line {
	s.newLine(*jumble)
	return &s.line
}
