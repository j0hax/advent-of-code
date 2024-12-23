package twentythree

import (
	"fmt"
	"io"
	"slices"
	"strings"
)

func PartOne(r io.Reader) int {
	graph := ParseGraph(r)

	cliques := graph.Cliques(3)

	// Search for cliques with a host that contains t
	candidates := 0
	for _, cl := range cliques {
		for _, c := range cl {
			if c[0] == 't' {
				candidates++
				break
			}
		}
	}

	return candidates
}

func PartTwo(r io.Reader) int {
	graph := ParseGraph(r)

	//nodesList := graph.Nodes()
	nodesList := make(map[string]struct{})
	for key := range graph {
		nodesList[key] = struct{}{}
	}

	maxLenClique := make(map[string]struct{})
	graph.BronKerbosch(make(map[string]struct{}), nodesList, make(map[string]struct{}), maxLenClique)

	maxLexNodes := make([]string, 0, len(maxLenClique))
	for key := range maxLenClique {
		fmt.Println(key)
		maxLexNodes = append(maxLexNodes, key)
	}

	slices.Sort(maxLexNodes)

	answer := strings.Join(maxLexNodes, ",")
	fmt.Println(answer)

	return 0
}
