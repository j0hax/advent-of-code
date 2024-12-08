package main

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"

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
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	//fmt.Printf("=== %s:\n", inputFile)
	fmt.Fprintf(w, "%s\t\n", inputFile)
	for i, p := range parts {
		file, err := os.Open(inputFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		result := p(file)

		//fmt.Printf("Part %d: %d\n", i+1, result)
		fmt.Fprintf(w, "Part %d\t%d\n", i+1, result)
	}
	w.Flush()
}

func main() {
	RunAll("./input1", one.PartOne, one.PartTwo)
	RunAll("./input2", two.PartOne, two.PartTwo)
	RunAll("./input3", three.PartOne, three.PartTwo)
	RunAll("./input4", four.PartOne, four.PartTwo)
	RunAll("./input5", five.PartOne, five.PartTwo)
	//RunAll("./input6", six.PartOne, six.PartTwo)
	RunAll("./input7", seven.PartOne, seven.PartTwo)
}
