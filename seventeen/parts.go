package seventeen

import (
	"fmt"
	"io"
	"slices"
)

func PartOne(r io.Reader) int {
	prog := ParseComputer(r)
	res := prog.Run()
	fmt.Println(res)

	// TODO: Return a string
	return 0
}

/*
func PartTwo(r io.Reader) int {
	prog := ParseComputer(r)

	// Create two queues
	aQ := []int{0}
	iQ := []int{15}

	var quines []int

	for len(aQ) > 0 {
		// pop queues
		a, aQ := aQ[0], aQ[1:]
		i, iQ := iQ[0], aQ[1:]

		for o := range 8 {
			test_a := (a << 3) + o
			target := slices.Clone(prog.Program)
		}

		res := prog.Run()
		fmt.Println(res)
	}

	// TODO: Return a string
	return 0
}
*/
