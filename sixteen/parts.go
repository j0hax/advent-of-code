package sixteen

import (
	"io"
)

func PartOne(r io.Reader) int {
	m := ParseMaze(r)

	return m.Solve()
}
