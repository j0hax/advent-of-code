package eight

import "io"

func PartOne(r io.Reader) int {
	smap := ParseMap(r)
	return smap.CountAntinodes(false)
}

func PartTwo(r io.Reader) int {
	smap := ParseMap(r)
	return smap.CountAntinodes(true)
}
