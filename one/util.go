package one

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"
)

/*
parseList opens the file passed to it, parses the first and second integer

	of each line and returns the sorted columns.
*/
func parseList(input io.Reader) (left []int, right []int) {
	scanner := bufio.NewScanner(input)

	left = make([]int, 0)
	right = make([]int, 0)

	/* NOTE: there are less verbose ways of reading a big file of ints into one
	   (or more) list(s). I deliberately chose an idiomatic way of taking the
	   first and second number of each line in the file.
	*/
	for scanner.Scan() {
		line := scanner.Text()

		nums := strings.Fields(line)

		l, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}

		r, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		left = append(left, l)
		right = append(right, r)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Sort
	slices.Sort(left)
	slices.Sort(right)

	return left, right
}
