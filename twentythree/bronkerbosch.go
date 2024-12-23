package twentythree

import (
	"maps"
)

func (g Graph[T]) BronKerbosch(R, P, X, max map[T]struct{}) {
	//fmt.Printf("%v %v %v\n", R, P, X)
	if len(P) == 0 && len(X) == 0 {
		if len(R) > len(max) {
			//max = maps.Clone(R)
			for v := range max {
				delete(max, v)
			}
			maps.Copy(max, R)
		}
		return
	}

	// for each vertex in P
	for v := range P {
		//fmt.Println(v)
		newR := maps.Clone(R)
		newR[v] = struct{}{}

		newP := make(map[T]struct{})
		newX := make(map[T]struct{})

		for _, n := range g[v] {
			//fmt.Println(n)
			if _, ok := P[n]; ok {
				newP[n] = struct{}{}
			}
			if _, ok := X[n]; ok {
				newX[n] = struct{}{}
			}
		}

		g.BronKerbosch(newR, newP, newX, max)

		// move V from P to X
		delete(P, v)
		X[v] = struct{}{}
	}
}
