package main

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"
	"time"

	"github.com/j0hax/aoc2024/five"
	"github.com/j0hax/aoc2024/four"
	"github.com/j0hax/aoc2024/one"
	"github.com/j0hax/aoc2024/seven"

	//"github.com/j0hax/aoc2024/six"
	"github.com/j0hax/aoc2024/three"
	"github.com/j0hax/aoc2024/two"
)

// Part is the function signature of one of two parts for a day.
// It takes an io.Reader which contains the input, and returns the integer
// solution
type Part func(r io.Reader) int

// RunParts executes several part functions with one input file.
func RunParts(w io.Writer, inputFile string, parts ...Part) {
	fmt.Fprintf(w, "%s\tSolution\tTime\n", inputFile)
	for i, p := range parts {
		file, err := os.Open(inputFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		start := time.Now()
		result := p(file)
		elapsed := time.Since(start)

		fmt.Fprintf(w, "Part %d\t%d\t%s\n", i+1, result, elapsed)
	}
}

func main() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer w.Flush()

	RunParts(w, "./input1", one.PartOne, one.PartTwo)
	RunParts(w, "./input2", two.PartOne, two.PartTwo)
	RunParts(w, "./input3", three.PartOne, three.PartTwo)
	RunParts(w, "./input4", four.PartOne, four.PartTwo)
	RunParts(w, "./input5", five.PartOne, five.PartTwo)
	//RunAll("./input6", six.PartOne, six.PartTwo)
	RunParts(w, "./input7", seven.PartOne, seven.PartTwo)
}
