package ent

type diagonal struct {
	line
}

/*func (s *diagonal) makeLine(jumble *[][]IdxPosition) *line {
	dst := make([][]IdxPosition, 0)
	copy(dst, *jumble)
	utl.Rotate(dst)
	utl.DiagonalView(dst)
	s.newLine(dst)
	return &s.line
}
*/
