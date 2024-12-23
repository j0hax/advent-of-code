package twenty

import (
	"io"
	"slices"
)

func PartOne(r io.Reader) int {
	m := ParseTrack(r)

	cheatList := m.CountCheats()

	cheatCounts := make(map[int]int)
	for _, v := range cheatList {
		cheatCounts[v]++
	}

	var keys []int
	for k := range cheatCounts {
		keys = append(keys, k)
	}

	slices.Sort(keys)

	goodCheat := 0

	for _, v := range keys {
		if v == 0 {
			continue
		}
		if v >= 100 {
			goodCheat += cheatCounts[v]
		}
	}

	return goodCheat
}

func PartTwo(r io.Reader) int {
	track := ParseTrack(r)
	return track.CountCheatsLen()
}
