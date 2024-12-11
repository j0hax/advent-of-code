package eleven

import (
	"io"
	"math"
	"strconv"
	"strings"
)

type Stone int

func (s Stone) Length() int {
	return int(math.Log10(float64(s))) + 1
}

func (s Stone) Split() (Stone, Stone) {
	str := strconv.Itoa(int(s))
	half := s.Length() / 2

	l := str[:half]
	r := str[half:]

	left, err := strconv.Atoi(l)
	if err != nil {
		panic(err)
	}

	right, err := strconv.Atoi(r)
	if err != nil {
		panic(err)
	}

	return Stone(left), Stone(right)
}

type state struct {
	stone Stone
	times int
}

var memoization = make(map[state]int)

func (stone Stone) Blink(times int) int {
	if times == 0 {
		return 1
	}

	if val, ok := memoization[state{stone, times}]; ok {
		return val
	}

	count := 0

	// Rule 1
	if stone == 0 {
		count = Stone(1).Blink(times - 1)
	} else if stone.Length()%2 == 0 {
		// Rule 2
		left, right := stone.Split()
		count = left.Blink(times-1) + right.Blink(times-1)
	} else {
		count = Stone(stone * 2024).Blink(times - 1)
	}

	memoization[state{stone, times}] = count

	return count
}

func ParseStones(r io.Reader) []Stone {
	var stones []Stone

	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		panic(err)
	}

	nums := strings.Fields(buf.String())

	for _, n := range nums {
		i, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}

		stones = append(stones, Stone(i))
	}

	return stones
}
