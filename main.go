package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/tabwriter"
	"time"

	"github.com/j0hax/aoc2024/eight"
	"github.com/j0hax/aoc2024/eleven"
	"github.com/j0hax/aoc2024/fifteen"
	"github.com/j0hax/aoc2024/five"
	"github.com/j0hax/aoc2024/four"
	"github.com/j0hax/aoc2024/fourteen"
	"github.com/j0hax/aoc2024/nine"
	"github.com/j0hax/aoc2024/nineteen"
	"github.com/j0hax/aoc2024/one"
	"github.com/j0hax/aoc2024/seven"
	"github.com/j0hax/aoc2024/six"
	"github.com/j0hax/aoc2024/ten"
	"github.com/j0hax/aoc2024/thirteen"
	"github.com/j0hax/aoc2024/three"
	"github.com/j0hax/aoc2024/twelve"
	"github.com/j0hax/aoc2024/twenty"
	"github.com/j0hax/aoc2024/twentyone"

	//"github.com/j0hax/aoc2024/twentyone"
	"github.com/j0hax/aoc2024/twentytwo"
	"github.com/j0hax/aoc2024/two"
)

// Part is the function signature of one of two parts for a day.
// It takes an io.Reader which contains the input, and returns the integer
// solution
type Part func(r io.Reader) int

// RunParts executes several part functions with one input file.
func RunParts(w io.Writer, inputFile string, parts ...Part) {
	fname := filepath.Base(inputFile)
	//fmt.Fprintf(w, "%s\t\t\t\n", fname)
	for i, p := range parts {
		file, err := os.Open(inputFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		start := time.Now()
		result := p(file)
		elapsed := time.Since(start).Round(time.Microsecond)

		fmt.Fprintf(w, "[%s]\tPart %d\t%d\t%s\n", fname, i+1, result, elapsed)
	}
}

func main() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer w.Flush()

	RunParts(w, "./inputs/input1", one.PartOne, one.PartTwo)
	RunParts(w, "./inputs/input2", two.PartOne, two.PartTwo)
	RunParts(w, "./inputs/input3", three.PartOne, three.PartTwo)
	RunParts(w, "./inputs/input4", four.PartOne, four.PartTwo)
	RunParts(w, "./inputs/input5", five.PartOne, five.PartTwo)
	RunParts(w, "./inputs/input6", six.PartOne) //six.PartTwo)
	RunParts(w, "./inputs/input7", seven.PartOne, seven.PartTwo)
	RunParts(w, "./inputs/input8", eight.PartOne, eight.PartTwo)
	RunParts(w, "./inputs/input9", nine.PartOne, nine.PartTwo)
	RunParts(w, "./inputs/input10", ten.PartOne, ten.PartTwo)
	RunParts(w, "./inputs/input11", eleven.PartOne, eleven.PartTwo)
	RunParts(w, "./inputs/input12", twelve.PartOne, twelve.PartTwo)
	RunParts(w, "./inputs/input13", thirteen.PartOne, thirteen.PartTwo)
	RunParts(w, "./inputs/input14", fourteen.PartOne, fourteen.PartTwo)
	RunParts(w, "./inputs/input15", fifteen.PartOne, fifteen.PartTwo)
	RunParts(w, "./inputs/input19", nineteen.PartOne, nineteen.PartTwo)
	RunParts(w, "./inputs/input20", twenty.PartOne, twenty.PartTwo)
	RunParts(w, "./inputs/input21", twentyone.PartOne, twentyone.PartTwo)
	RunParts(w, "./inputs/input22", twentytwo.PartOne, twentytwo.PartTwo)
}
