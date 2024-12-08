package five

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"
)

type Rule map[int][]int

type Update []int

type Protocol struct {
	rules   Rule
	updates []Update
}

// Correct cycles through the page numbers of an update.
// If a page is determined to be out of order, false is returned.
func (u Update) Correct(rules Rule) bool {
	seen := make(map[int]bool)
	// Cycle through each page of an update,
	// then go through and ensure that each
	for _, page := range u {
		previous := rules[page]
		for _, v := range previous {
			// Check if page number is contained in previous
			if _, ok := seen[v]; ok {
				return false
			}
		}
		// If we have made it to the end, place a value in seen.
		seen[page] = true
	}

	return true
}

// Reorder reorders the numbers in an update in place according to the given
// rules
func (u Update) Reorder(rules Rule) {
	slices.SortFunc(u, func(a, b int) int {

		r := rules[a]
		c := slices.Contains(r, b)

		if c {
			return -1
		} else {
			return 1
		}
	})
}

/*
FixIncorrect is part two of day five. It loops through each update,

verifying that the order is correct. If the order of an update is incorrect, it
is reordered. The sum of middle numbers of each corrected update is returned.
*/
func (p *Protocol) FixIncorrect() int {
	sum := 0

	for _, u := range p.updates {
		//results = append(results, u.Correct(p.rules))
		c := u.Correct(p.rules)

		if !c {
			u.Reorder(p.rules)
			l := len(u) / 2
			sum += u[l]
		}
	}

	return sum
}

// Correct checks if each page in the update is in an order according to rules
// and returns the sum of the middle part.
func (p *Protocol) Correct() int {
	//var results []bool
	sum := 0
	for _, u := range p.updates {
		//results = append(results, u.Correct(p.rules))
		c := u.Correct(p.rules)
		//fmt.Printf("Row %d correct: %v\n", i, c)

		if c {
			l := len(u) / 2
			sum += u[l]
		}
	}

	return sum
}

// splitToInt splits string s according to sep, then converts each field to an
// integer
func splitToInt(s, sep string) []int {
	var list []int
	fields := strings.Split(s, sep)

	for _, num := range fields {
		n, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}

		list = append(list, n)
	}

	return list
}

// ReadProtocol reads a file from disk and parses the contents to a Protocol
func ReadProtocol(r io.Reader) *Protocol {
	p := Protocol{
		rules: make(map[int][]int),
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		// Determine if the line is a rule or an update
		line := scanner.Text()
		if strings.ContainsRune(line, '|') {
			// Split the first and last part, then add to the map
			n := splitToInt(line, "|")
			p.rules[n[0]] = append(p.rules[n[0]], n[1])
		} else if strings.ContainsRune(line, ',') {
			u := splitToInt(line, ",")
			p.updates = append(p.updates, u)
		}
	}

	return &p
}
