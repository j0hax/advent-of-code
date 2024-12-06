package six

import (
	"fmt"
	"io"
)

func PartOne(r io.Reader) int {
	world := Load(r)

	for world.Step() != oob {

	}

	fmt.Println(world)

	return world.Count(visited)
}
