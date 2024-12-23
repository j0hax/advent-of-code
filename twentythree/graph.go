package twentythree

import (
	"bufio"
	"cmp"
	"fmt"
	"io"
	"slices"
	"strings"
)

// Graph represents an undirected Graph. It is implemented as an adjacency list.
type Graph[T cmp.Ordered] map[T][]T

// Edge adds an undirected edge to the graph
func (g Graph[T]) Edge(a, b T) {
	g[a] = append(g[a], b)
	g[b] = append(g[b], a)
}

// name is used to create a reproducible name to eliminate duplicate cliques
// using a hashmap
func name[T cmp.Ordered](members []T) string {
	slices.Sort(members)
	var sb strings.Builder
	for _, ob := range members {
		sb.WriteString(fmt.Sprintf("%v", ob))
	}
	return sb.String()
}

func (g Graph[T]) cliqrec(k int, knownCliques map[string][]T, members ...T) {
	last := members[len(members)-1]

	if k == len(members) {
		if slices.Contains(g[last], members[0]) {
			// The clique is a full cycle, now we check if it exists
			n := name(members)
			if _, ok := knownCliques[n]; !ok {
				knownCliques[n] = members
			}

		}
		return
	}

	for _, val := range g[last] {
		g.cliqrec(k, knownCliques, append(members, val)...)
	}

	return
}

// Cliques searches through the graph for Cliques of size k
func (g Graph[T]) Cliques(k int) (cliques [][]T) {
	knownCliques := make(map[string][]T)

	for key := range g {
		g.cliqrec(k, knownCliques, key)
	}

	for _, v := range knownCliques {
		cliques = append(cliques, v)
	}

	return
}

func (g Graph[T]) Nodes() []T {
	nodesList := make([]T, 0, len(g))
	for v := range g {
		nodesList = append(nodesList, v)
	}

	return nodesList
}

func ParseGraph(r io.Reader) (g Graph[string]) {
	g = make(Graph[string])

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		temp := strings.Split(scanner.Text(), "-")
		g.Edge(temp[0], temp[1])
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return
}
