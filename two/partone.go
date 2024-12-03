package two

import "os"
import "bufio"
import "strings"
import "fmt"
import "strconv"

type Report []int

// atoiSlice accepts a slice of stringified integers and returns a slice of
// regular integers.
func atoiSlice(fields []string) ([]int, error) {
	var intFields []int

	for i := 0; i < len(fields); i++ {
		rep, err := strconv.Atoi(fields[i])
		if err != nil {
			return intFields, err
		}
		intFields = append(intFields, rep)
	}

	return intFields, nil
}

func ParseRecords(inputFile string) []Report {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var r []Report

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		fields, err := atoiSlice(strings.Fields(line))
		if err != nil {
			panic(err)
		}

		r = append(r, fields)
	}

	return r
}

// ShouldInc checks whether the first two non-equal values are increasing or
// decreasing in order.
func (r Report) ShouldInc() bool {
	//first := r[0]
	for i := 1; i < len(r); i++ {
		if r[i-1] != r[i] {
			return r[i-1] < r[i]
		}
	}

	return false
}

/*
Safe checks if the records are "safe," i.e. steadily increasing or decreasing.

NOTE: you really have to read carefully. I got big mad because I only realized
after quite some time that my algorithm was almost correct, but I permitted two
of the same numbers after another (diff = 0) when that was not safe...
*/
func (r Report) Safe() bool {
	// Check if increasing or decreasing
	inc := r.ShouldInc()

	for i := 1; i < len(r); i++ {
		diff := r[i] - r[i-1]

		if inc && (diff > 3 || diff < 1) {
			//fmt.Printf("%v NOT safe: increasing=%v but diff=%d\n", r, inc, diff)
			return false
		}

		if !inc && (diff < -3 || diff > -1) {
			//fmt.Printf("%v NOT safe: increasing=%v but diff=%d\n", r, inc, diff)
			return false
		}
	}

	//fmt.Printf("%v appears safe\n", r)
	return true
}

func PartOne() {
	records := ParseRecords("./input2")

	safeCnt := 0

	for _, rep := range records {
		if rep.Safe() {
			safeCnt++
		}
	}

	fmt.Printf("Solution for Part One: %d\n", safeCnt)
}
