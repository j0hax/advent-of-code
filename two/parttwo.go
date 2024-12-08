package two

import (
	"io"
)

// removeAtIndex returns a new Report with the items deleted at index i.
// Note that this function returns a NEW slice, not the modified slice.
func removeAtIndex(s Report, i int) Report {
	//return append(s[:i], s[i+1:]...)
	r := make([]int, 0, len(s)-1)
	r = append(r, s[:i]...)
	return append(r, s[i+1:]...)
}

/*
DampenedSafe is the Brute-Force dampen search: if a record is not safe, any

	of its elements is removed once to see if it is then safe.

	NOTE: I likely had more elegant solution for this in linear time at some
	point... and assumed the solution was wrong, because I was fighting with the
	fact that slices are references.
*/
func (r Report) DampenedSafe() bool {
	// Try one
	if r.Safe() {
		return true
	}

	// Brute-Force: Remove any individual level and check if it is safe now
	for i := 0; i < len(r); i++ {
		cut := removeAtIndex(r, i)
		//fmt.Printf("@%d %v -> %v\n", i, r, cut)
		if cut.Safe() {
			return true
		}
	}

	return false
}

func PartTwo(r io.Reader) int {
	records := ParseRecords(r)
	safeCnt := 0

	for _, rep := range records {
		if rep.DampenedSafe() {
			safeCnt++
		}
	}

	return safeCnt
}
