package main

import (
	"strings"
	"testing"
)

func TestSum(t *testing.T) {
	input := `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`

	r := strings.NewReader(input)
	s := SumDocument(r)

	expected := 142

	if s != expected {
		t.Fatalf(`Sum() from example should be %v, got %v`, expected, s)
	}
}
