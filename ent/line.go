package ent

import "strings"

type line struct {
	IdxPosition
	line string
	idx  []IdxPosition
}

func (s line) find(str string) IdxPosition {
	index := strings.Index(s.line, str)
	if index == -1 {
		return s.notFound()
	}
	return s.idx[index]
}

func (s *line) newLine(tj [][]IdxPosition) {
	for _, x := range tj {
		for _, y := range x {
			s.line = s.line + y.letter
			s.idx = append(s.idx, y)
		}
	}
	s.line = strings.ToUpper(s.line)
}
