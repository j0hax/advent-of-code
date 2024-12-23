package twentytwo

import (
	"bufio"
	"io"
	"strconv"
)

type SecretNumber int

// MixPrune first performs a bitwise operation of m, then calculates the
// modulo with 16777216
func (s *SecretNumber) MixPrune(m SecretNumber) {
	*s ^= m
	*s %= 16777216
}

// Next calculates the next pseudorandom number of s
func (s SecretNumber) Next() SecretNumber {
	s.MixPrune(s * 64)
	s.MixPrune(s / 32)
	s.MixPrune(s * 2048)

	return s
}

func (s SecretNumber) Price() int {
	return int(s) % 10
}

func ParseNumbers(r io.Reader) []SecretNumber {
	var numbers []SecretNumber

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, SecretNumber(n))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return numbers
}
