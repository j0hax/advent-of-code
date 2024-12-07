package seven

import (
	"io"
)

func PartOne(r io.Reader) int {
	total := 0

	eqs := ParseEquations(r)
	for _, e := range eqs {
		if e.CountSolutions(false) > 0 {
			total += e.total
		}
	}

	return total
}

func PartTwo(r io.Reader) int {
	total := 0

	eqs := ParseEquations(r)
	for _, e := range eqs {
		if e.CountSolutions(true) > 0 {
			total += e.total
		}
	}

	return total
}
