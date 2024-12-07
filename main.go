package main

import (
	"fmt"
	"io"
	"os"

	"github.com/j0hax/aoc2024/five"
	"github.com/j0hax/aoc2024/four"
	"github.com/j0hax/aoc2024/one"
	"github.com/j0hax/aoc2024/seven"
	//"github.com/j0hax/aoc2024/six"
	"github.com/j0hax/aoc2024/three"
	"github.com/j0hax/aoc2024/two"
)

type Part func(r io.Reader) int

func RunAll(inputFile string, parts ...Part) {
	for i, p := range parts {
		file, err := os.Open(inputFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		result := p(file)

		fmt.Printf("Solution for Part %d: %d\n", i, result)
	}
}

func main() {
	//one.PartOne()
	//one.PartTwo()
	RunAll("./input1", one.PartOne, one.PartTwo)

	two.PartOne()
	two.PartTwo()

	three.PartOne()
	three.PartTwo()

	four.PartOne()
	four.PartTwo()

	five.PartOne()
	five.PartTwo()

	//RunAll("./input6", six.PartOne, six.PartTwo)

	RunAll("./input7", seven.PartOne)
}
