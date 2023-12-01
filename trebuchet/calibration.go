package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

// Returns the sum of the first and last digits contained in line
func SumLine(line string) int {
	first := 0
	last := 0

	// Convert string to rune array
	dat := []rune(line)

	// Forward search
	for i := 0; i < len(dat); i++ {
		if unicode.IsDigit(dat[i]) {
			last = int(dat[i] - '0')
		}
	}

	// Backward search
	for i := len(dat) - 1; i >= 0; i-- {
		if unicode.IsDigit(dat[i]) {
			first = int(dat[i] - '0')
		}
	}

	combine := first*10 + last

	log.Printf("%v + %v = %v\n", first, last, combine)

	return combine
}

func SumDocument(doc io.Reader) int {
	sum := 0

	scanner := bufio.NewScanner(doc)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		sum += SumLine(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Panic(err)
	}

	return sum
}

func main() {
	result := SumDocument(os.Stdin)

	fmt.Printf("Result:\t%v\n", result)
}
