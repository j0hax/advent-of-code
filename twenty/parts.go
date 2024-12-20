package twenty

import (
	"fmt"
	"io"
	"slices"
)

func PartOne(r io.Reader) int {
	m := ParseTrack(r)
	fmt.Println(m)

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
		fmt.Printf("There are %d cheats that save %d picoseconds\n", cheatCounts[v], v)
		if v >= 100 {
			goodCheat += cheatCounts[v]
		}
	}

	return goodCheat
}
