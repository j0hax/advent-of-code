package five

import "fmt"

func PartOne() {
	prot := ReadProtocol("./input5")
	s := prot.Correct()
	fmt.Printf("Solution to part one: %d\n", s)
}

func PartTwo() {
	prot := ReadProtocol("./input5")
	s := prot.FixIncorrect()
	fmt.Printf("Solution to part two: %d\n", s)
}
