package main

import (
	"strings"
	"testing"
)

func TestSumBasic(t *testing.T) {
	input := `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`

	r := strings.NewReader(input)
	s := SumDocument(r, false)

	expected := 142

	if s != expected {
		t.Fatalf(`Sum() from example should be %v, got %v`, expected, s)
	}
}

func TestSumSpelled(t *testing.T) {
	input := `two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`

	r := strings.NewReader(input)
	s := SumDocument(r, true)

	expected := 281

	if s != expected {
		t.Fatalf(`Sum() from example should be %v, got %v`, expected, s)
	}
}
