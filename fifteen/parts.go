package fifteen

import (
	"fmt"
	"io"
)

func PartOne(r io.Reader) int {
	w := ParseWareHouse(r)

	fmt.Println(w)

	for len(w.Moves) > 0 {
		w.Step()
		//fmt.Println(w)
	}

	return w.SumBoxes()
}
