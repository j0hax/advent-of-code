package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

// Map containing spelled digits.
var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

/*
The big problem with spelled out numbers:

you can't simply replace the word with the digit in the string:
e.g. "oneight" would be "1ight," which would evaluate to 11 and not 18.

As such, a simple forward parser had to be implemented.
*/
func parseNum(dat []rune, index int, includeWords bool) int {
	// Part one: find a digit
	if unicode.IsDigit(dat[index]) {
		return int(dat[index] - '0')
	}

	// Part two: find a letter, check if it matches what we know
	if includeWords {
		for k, v := range digits {
			// quick oob check
			indexEnd := index + len(k)
			if indexEnd > len(dat) {
				continue
			}

			// We know the key. Now we compare the key to the next n letters of data
			nextTokens := string(dat[index:indexEnd])
			if k == nextTokens {
				log.Printf("Match: %s means %d\n", nextTokens, v)
				return v
			}
		}
	}

	return -1
}

// Returns the sum of the first and last digits contained in line
func SumLine(line string, includeWords bool) int {
	dat := []rune(line)
	first := 0
	last := 0

	// Forward search
	for i := 0; i < len(dat); i++ {
		res := parseNum(dat, i, includeWords)
		if res > 0 {
			first = res
			break
		}
	}

	// Backward search
	for i := len(dat) - 1; i >= 0; i-- {
		res := parseNum(dat, i, includeWords)
		if res > 0 {
			last = res
			break
		}
	}

	result := first*10 + last

	log.Printf("%v => (%d, %d) = %d\n", line, first, last, result)

	return result
}

func SumDocument(doc io.Reader, toDecimal bool) int {
	sum := 0

	scanner := bufio.NewScanner(doc)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		sum += SumLine(scanner.Text(), toDecimal)
	}
	if err := scanner.Err(); err != nil {
		log.Panic(err)
	}

	return sum
}

func main() {
	spelled := flag.Bool("spelled", false, "Include spelled words")
	flag.Parse()

	result := SumDocument(os.Stdin, *spelled)

	fmt.Printf("Result:\t%v\n", result)
}
