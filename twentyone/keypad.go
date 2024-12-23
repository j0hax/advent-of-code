/*
The code for day 21 was adapted with permission from the following repository:
https://github.com/arjunpathak072/aoc-2024/blob/main/day-21/main.go
*/
package twentyone

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"slices"
	"strconv"
)

type Position struct {
	r, c int
}

func (p Position) Add(n Position) Position {
	return Position{
		p.r + n.r,
		p.c + n.c,
	}
}

var dirMap = map[rune]Position{
	'^': {-1, 0},
	'v': {1, 0},
	'>': {0, 1},
	'<': {0, -1},
}

var numericKeypad = map[rune]Position{
	'7': {0, 0}, '8': {0, 1}, '9': {0, 2},
	'4': {1, 0}, '5': {1, 1}, '6': {1, 2},
	'1': {2, 0}, '2': {2, 1}, '3': {2, 2},
	'0': {3, 1}, 'A': {3, 2},
}

var directionKeypad = map[rune]Position{
	'^': {0, 1}, 'A': {0, 2},
	'<': {1, 0}, 'v': {1, 1}, '>': {1, 2},
}

var revDirectionKeypad = getReverseMap(directionKeypad)
var revNumericKeypad = getReverseMap(numericKeypad)

var pairsMinDistanceCache = make(map[string]int)
var pathsCache = make(map[string][]Code)

// Code represents a sequences of characters input into a keypad
type Code []rune

// Coeff returns the code's first n-1 numbers
func (c Code) Coeff() int {
	i := len(c) - 1
	coeff, err := strconv.Atoi(string(c[:i]))
	if err != nil {
		panic(err)
	}
	return coeff
}

// Prepend adds r runes to the front of the sequence (in place).
func (c *Code) Prepend(r ...rune) {
	*c = append(r, *c...)
}

// Append adds r runes to the back of the sequence (in place).
func (c *Code) Append(r ...rune) {
	*c = append(*c, r...)
}

// ParseCodes parses a sequence of codes from an input source
func ParseCodes(r io.Reader) (codes []Code) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		newCode := Code(scanner.Text())
		codes = append(codes, newCode)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return
}

// Solve calculates the number of required moves *depth* robots.
func (input Code) Solve(depth int) (res int) {
	coeff := input.Coeff()

	// Calculate cost
	input.Prepend('A')
	for i := 0; i < len(input)-1; i++ {
		currPairCost := GetPairCost(rune(input[i]), rune(input[i+1]), numericKeypad, revNumericKeypad, depth)
		res += currPairCost
	}

	res *= coeff
	return
}

// GetPairCost calculates the cost, i.e. moves required, from key a to b
func GetPairCost(a, b rune, charToIndex map[rune]Position, indexToChar map[Position]rune, depth int) int {
	keypadCode := 'd'
	if _, ok := charToIndex['0']; ok {
		keypadCode = 'n'
	}
	key := fmt.Sprintf("%c%c%c%d", a, b, keypadCode, depth)

	if dist, ok := pairsMinDistanceCache[key]; ok {
		return dist
	}

	if depth == 0 {
		minLen := math.MaxInt
		for _, path := range GetAllPaths(a, b, directionKeypad, revDirectionKeypad) {
			minLen = min(minLen, len(path))
		}
		return minLen
	}

	allPaths := GetAllPaths(a, b, charToIndex, indexToChar)
	minCost := math.MaxInt

	for _, path := range allPaths {
		path.Prepend('A')
		var currCost int

		for i := 0; i < len(path)-1; i++ {
			currCost += GetPairCost(rune(path[i]), rune(path[i+1]), directionKeypad, revDirectionKeypad, depth-1)
		}
		minCost = min(minCost, currCost)
	}

	pairsMinDistanceCache[key] = minCost
	return minCost
}

// GetAllPaths uses a depth-first search to calculate all possible paths from
// key a to b
func GetAllPaths(a, b rune, charToIndex map[rune]Position, indexToChar map[Position]rune) (allPaths []Code) {
	key := fmt.Sprintf("%c %c", a, b)
	if paths, ok := pathsCache[key]; ok {
		return paths
	}
	DFS(charToIndex[a], charToIndex[b], []rune{}, charToIndex, indexToChar, make(map[Position]bool), &allPaths)
	pathsCache[key] = allPaths
	return
}

func DFS(curr, end Position, path Code, charToIndex map[rune]Position, indexToChar map[Position]rune, visited map[Position]bool, allPaths *[]Code) {
	if curr == end {
		path.Append('A')
		*allPaths = append(*allPaths, path)
		return
	}
	visited[curr] = true
	for char, dir := range dirMap {
		nIdx := curr.Add(dir)
		if _, ok := indexToChar[nIdx]; ok && !visited[nIdx] {
			newPath := slices.Clone(path)
			DFS(nIdx, end, append(newPath, char), charToIndex, indexToChar, visited, allPaths)
		}
	}
	visited[curr] = false
}

func getReverseMap(m map[rune]Position) (w map[Position]rune) {
	w = make(map[Position]rune)
	for r, i := range m {
		w[i] = r
	}
	return
}
