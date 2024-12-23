package eighteen

import (
	"bytes"
	"fmt"
	"io"
)

func PartOne(r io.Reader) int {
	m := ParseRAM(r, 1024)
	return m.Solve()
}

func PartTwo(r io.Reader) int {
	buf, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	size := 1024
	for {
		m := ParseRAM(bytes.NewReader(buf), size)
		if m.Solve() == -1 {
			fmt.Println(m.lastLoc)
			return 0
		}
		size++
	}
}
