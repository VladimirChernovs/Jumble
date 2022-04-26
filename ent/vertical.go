package ent

import (
	"Jumble/utl"
)

type vertical struct {
	line
}

func (s *vertical) makeLine(jumble *[][]IdxPosition) *line {
	tj := utl.TransposeSlices(*jumble)
	s.newLine(tj)
	return &s.line
}
