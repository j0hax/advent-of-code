package five

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule map[int][]int

type Update []int

type Protocol struct {
	rules   Rule
	updates []Update
}

func (u Update) Correct(rules Rule) bool {
	seen := make(map[int]bool)
	for _, page := range u {
		previous := rules[page]
		for _, p := range previous {
			if _, ok := seen[p]; ok {
				return false
			}
		}
		seen[page] = true
	}

	return true
}

// Correct checks if each page in the update is in an order according to rules
// and returns the sum of the middle part.
func (p *Protocol) Correct() int {
	//var results []bool
	sum := 0
	for i, u := range p.updates {
		//results = append(results, u.Correct(p.rules))
		c := u.Correct(p.rules)
		fmt.Printf("Row %d correct: %v\n", i, c)

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

func ReadProtocol(filename string) *Protocol {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := Protocol{
		rules: make(map[int][]int),
	}

	scanner := bufio.NewScanner(file)
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
