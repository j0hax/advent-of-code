package thirteen

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Pair struct {
	X, Y int
}

func (p Pair) Add(n Pair) Pair {
	return Pair{p.X + n.X, p.Y + n.Y}
}

func (p Pair) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func (p Pair) Equal(n Pair) bool {
	return p.X == n.X && p.Y == n.Y
}

type Machine struct {
	A, B, Prize Pair
}

func (m Machine) String() string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "Button A: X+%d, Y+%d\n", m.A.X, m.A.Y)
	fmt.Fprintf(&sb, "Button B: X+%d, Y+%d\n", m.B.X, m.B.Y)
	fmt.Fprintf(&sb, "Prize: X=%d, Y=%d\n", m.Prize.X, m.Prize.Y)

	return sb.String()
}

func (m Machine) Win() int {
	d := (m.A.X*m.B.Y - m.A.Y*m.B.X)

	a := (m.Prize.X*m.B.Y - m.B.X*m.Prize.Y) / d
	b := (m.A.X*m.Prize.Y - m.Prize.X*m.A.Y) / d

	// Check if we get the target
	if a >= 0 && b >= 0 && m.A.X*a+m.B.X*b == m.Prize.X && m.A.Y*a+m.B.Y*b == m.Prize.Y {
		return a*3 + b
	}

	return 0
}

func (m Machine) AdjWin() int {
	m.Prize.X += 10000000000000
	m.Prize.Y += 10000000000000

	return m.Win()
}

func ParseMachines(r io.Reader) []Machine {
	var machines []Machine

	re := regexp.MustCompile(`\d+`)

	scanner := bufio.NewScanner(r)
	currMach := Machine{}

	// Always read two numbers
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		nums := re.FindAllString(line, 2)

		n1, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}

		n2, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		p := Pair{n1, n2}

		// Check where to put data
		if currMach.A.X == 0 {
			currMach.A = p
		} else if currMach.B.X == 0 {
			currMach.B = p
		} else if currMach.Prize.X == 0 {
			currMach.Prize = p

			// Last field for a machine, so copy and create a new one
			machines = append(machines, currMach)
			currMach = Machine{}
		}
	}

	return machines
}
