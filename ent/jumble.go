package ent

import "strings"

type Jumble struct {
	matrix     [][]IdxPosition
	horizontal *line
	vertical   *line
	//diagonals  *line
}

func (j Jumble) Find(str string) (int, int) {
	return j.findIdx(str).GetIdx()
}

func (j Jumble) findIdx(s string) IdxPosition {
	str := strings.ToUpper(s)
	fs := make([]func(string) IdxPosition, 0)
	fs = append(fs, j.vertical.find)
	fs = append(fs, j.horizontal.find)
	return searchData(str, fs)
}

func NewJumble(jumble [][]string) *Jumble {
	j := new(Jumble)
	makeJumble(jumble, j)

	h := make(chan *line)
	v := make(chan *line)
	//d := make(chan *line)

	go func() { h <- new(horizontal).makeLine(&j.matrix) }()
	go func() { v <- new(vertical).makeLine(&j.matrix) }()
	//go func() { d <- new(diagonal).makeLine(&j.matrix) }()
	j.horizontal = <-h
	j.vertical = <-v

	return j
}

func makeJumble(jumble [][]string, j *Jumble) {
	for i, lC := range jumble {
		mx := make([]IdxPosition, 0)
		for j, _ := range lC {
			position := IdxPosition{letter: jumble[i][j], i: i, j: j}
			mx = append(mx, position)
		}
		j.matrix = append(j.matrix, mx)
	}
}

func searchData(s string, searchers []func(string) IdxPosition) IdxPosition {
	done := make(chan struct{})
	result := make(chan IdxPosition)
	for _, searcher := range searchers {
		go func(searcher func(string) IdxPosition) {
			select {
			case result <- searcher(s):
			case <-done:
			}
		}(searcher)
	}
	var rzt IdxPosition
	for r := range result {
		if !(r.i == -1 && r.j == -1) {
			close(done)
			return r
		}
		rzt = r
	}
	return rzt
}
