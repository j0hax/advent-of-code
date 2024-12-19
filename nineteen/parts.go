package nineteen

import (
	"io"
)

func PartOne(r io.Reader) int {
	m := ParseDesigns(r)
	return m.CountPossible()
}
