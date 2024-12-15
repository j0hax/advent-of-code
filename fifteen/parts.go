package fifteen

import (
	"io"
)

func PartOne(r io.Reader) int {
	w := ParseWareHouse(r, false)

	for len(w.Moves) > 0 {
		w.Step()
	}

	return w.SumBoxes()
}

func PartTwo(r io.Reader) int {
	w := ParseWareHouse(r, true)

	for len(w.Moves) > 0 {
		w.WideStep()
	}

	return w.SumBoxes()
}
