package five

import "io"

func PartOne(r io.Reader) int {
	prot := ReadProtocol(r)
	s := prot.Correct()
	return s
}

func PartTwo(r io.Reader) int {
	prot := ReadProtocol(r)
	s := prot.FixIncorrect()
	return s
}
