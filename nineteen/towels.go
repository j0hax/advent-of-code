package nineteen

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strings"
)

type Color rune

const (
	White Color = 'w'
	Blue        = 'u'
	Black       = 'b'
	Red         = 'r'
	Green       = 'g'
)

func (c Color) String() string {
	return fmt.Sprintf("%c", c)
}

type Towel []Color

func (t Towel) String() string {
	var sb strings.Builder
	for _, c := range t {
		sb.WriteRune(rune(c))
	}
	return sb.String()
}

type Designs struct {
	Available []Towel
	Goals     [][]Color
}

func CanMake(memo map[string]bool, goal []Color, avail []Towel) bool {
	if len(goal) == 0 {
		return true
	}

	if v, ok := memo[Towel(goal).String()]; ok {
		return v
	}

	// Check if next color can be used with any of the available
	for _, next := range avail {
		if len(next) <= len(goal) {
			head := goal[:len(next)]
			tail := goal[len(next):]

			// Check if each val of next satisfies goal
			if slices.Equal(next, head) {
				if CanMake(memo, tail, avail) {
					memo[Towel(tail).String()] = true
					return true
				}
			}
		}
	}

	memo[Towel(goal).String()] = false
	return false
}

func (d *Designs) CountPossible() int {
	count := 0
	for _, g := range d.Goals {
		fmt.Println(g)
		memo := make(map[string]bool)
		if CanMake(memo, g, d.Available) {
			count++
		}
	}

	return count
}

func ParseAvailTowels(s string) []Towel {
	var towels []Towel
	for _, t := range strings.Split(s, ", ") {
		towels = append(towels, []Color(t))
	}

	return towels
}

func ParseDesigns(r io.Reader) *Designs {
	scanner := bufio.NewScanner(r)

	var d Designs

	// Scan first line
	scanner.Scan()
	d.Available = ParseAvailTowels(scanner.Text())

	// Scan remaining
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		d.Goals = append(d.Goals, []Color(line))
	}

	return &d
}
