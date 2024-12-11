package eleven

import (
	"io"
)

func PartOne(r io.Reader) int {
	stones := ParseStones(r)

	sum := 0
	for _, s := range stones {
		sum += s.Blink(25)
	}

	return sum
}

func PartTwo(r io.Reader) int {
	stones := ParseStones(r)

	sum := 0
	for _, s := range stones {
		sum += s.Blink(75)
	}

	return sum
}
