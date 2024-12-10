package ten

import (
	"io"
)

func PartOne(r io.Reader) int {
	m := ParseMap(r)
	return m.ScoreTrails()
}

func PartTwo(r io.Reader) int {
	m := ParseMap(r)
	return m.GradeTrails()
}
