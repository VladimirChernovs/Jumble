package ent

type IdxPosition struct {
	letter string
	i      int
	j      int
}

func (s IdxPosition) GetIdx() (int, int) {
	return s.i, s.j
}

func (s IdxPosition) notFound() IdxPosition {
	return IdxPosition{"", -1, -1}
}

func (s IdxPosition) isFound() bool {
	return s.i == -1 && s.j == -1
}
