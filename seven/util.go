package seven

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Equation represents the total (left hand side) and a number of operands,
// without their respective operations.
type Equation struct {
	total    int
	operands []int
}

// Process replaces the first two items of the slice s with the result of the
// operator function.
func process(s []int, operator func(a, b int) int) []int {
	result := operator(s[0], s[1])
	return append([]int{result}, s[2:]...)
}

/*
	CountSolutions recursively operates through each operand and counts the

number of possible solutions for the equation's total. The method returns zero
if the total can not be formed by any combination of right-to-left operators.

NOTE: For day 7 part one, all we really had to do was determine if the equation
can be formed (not the count)... reading comprehension helps, folks ;)
*/
func (eq *Equation) CountSolutions(concatenate bool) int {
	if len(eq.operands) == 1 {
		if eq.total == eq.operands[0] {
			return 1
		} else {
			return 0
		}
	}

	taskList := []Equation{
		{
			total:    eq.total,
			operands: process(eq.operands, func(a, b int) int { return a * b }),
		},
		{
			total:    eq.total,
			operands: process(eq.operands, func(a, b int) int { return a + b }),
		},
	}

	if concatenate {
		taskList = append(taskList, Equation{
			total: eq.total,
			operands: process(eq.operands,
				func(a, b int) int {
					numStr := fmt.Sprintf("%d%d", a, b)
					num, err := strconv.Atoi(numStr)
					if err != nil {
						panic(err)
					}

					return num
				}),
		})
	}

	count := 0
	for _, task := range taskList {
		count += task.CountSolutions(concatenate)
	}

	return count
}

// ParseEquations reads the input stream and returns a list of equations.
func ParseEquations(r io.Reader) []Equation {
	var eqList []Equation

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")

		var eq Equation

		total, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		eq.total = total

		for _, s := range strings.Fields(parts[1]) {
			num, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}

			eq.operands = append(eq.operands, num)
		}

		eqList = append(eqList, eq)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return eqList
}
