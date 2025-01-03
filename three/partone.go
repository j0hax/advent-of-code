package three

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// parse accepts a string in the form of `mul(a,b)` and returns the product of
// a * b. If an error occurs, 0 and err is returned.
func parse(mul string) (int, error) {
	field := strings.Trim(mul, "mul()")
	nums := strings.Split(field, ",")
	a, err := strconv.Atoi(nums[0])
	if err != nil {
		return 0, err
	}

	b, err := strconv.Atoi(nums[1])
	if err != nil {
		return 0, err
	}

	return (a * b), nil
}

func PartOne(r io.Reader) int {
	solution := 0

	// The expression for mul:
	// match mul(x,x), with x containing between 1 to 3 digits.
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			s, err := parse(match)
			if err != nil {
				panic(err)
			}

			solution += s
		}
	}

	return solution
}

func PartTwo(r io.Reader) int {
	enabled := true
	solution := 0

	// Like above, but now with three items:
	// Match don't() or mul(x,y) or do()
	re := regexp.MustCompile(`(don't\(\)|mul\(\d{1,3},\d{1,3}\)|do\(\))`)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			if match == "don't()" {
				enabled = false
			} else if match == "do()" {
				enabled = true
			} else if enabled {
				s, err := parse(match)
				if err != nil {
					panic(err)
				}
				solution += s
			}
		}
	}
	return solution
}
