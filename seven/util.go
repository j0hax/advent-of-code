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

/*
	CountSolutions recursively operates through each operand and counts the

number of possible solutions for the equation's total. The method returns zero
if the total can not be formed by any combination of right-to-left operators.

NOTE: For day 7 part one, all we really had to do was determine if the equation
can be formed (not the count)... reading comprehension helps, folks ;)
*/
func (eq *Equation) CountSolutions() int {
	if len(eq.operands) == 1 {
		if eq.total == eq.operands[0] {
			return 1
		} else {
			return 0
		}
	}

	multResult := eq.operands[0] * eq.operands[1]
	multRest := append([]int{multResult}, eq.operands[2:]...)

	addResult := eq.operands[0] + eq.operands[1]
	addRest := append([]int{addResult}, eq.operands[2:]...)

	fmt.Printf("%v -> %v\n", eq.operands, addRest)

	// div
	recDiv := Equation{
		total:    eq.total,
		operands: multRest,
	}

	// minus
	recSub := Equation{
		total:    eq.total,
		operands: addRest,
	}

	return recDiv.CountSolutions() + recSub.CountSolutions()
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
