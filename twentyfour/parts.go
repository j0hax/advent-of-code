package twentyfour

import (
	"fmt"
	"io"
	"strconv"
)

func PartOne(r io.Reader) int {
	c := ParseCircuit(r)
	c.Simulate()

	result := 0

	for k, v := range c.inputs {
		if k[0] == 'z' {
			if v {
				num := string(k[1:])
				n, err := strconv.Atoi(num)
				if err != nil {
					panic(err)
				}
				result |= 1 << n
			}
		}
	}

	fmt.Println(result)

	return result
}
