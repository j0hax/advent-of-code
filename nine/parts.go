package nine

import (
	"io"
)

func PartOne(r io.Reader) int {
	dm := ParseMap(r)
	dm.Compact()
	return dm.CheckSum()
}

func PartTwo(r io.Reader) int {
	dm := ParseMap(r)
	dm.Defrag()
	return dm.CheckSum()
}
