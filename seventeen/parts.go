package seventeen

import (
	"fmt"
	"io"
)

func PartOne(r io.Reader) int {
	prog := ParseProg(r)
	res := prog.Run()
	fmt.Println(res)

	// TODO: Return a string
	return 0
}
