package graph

import "fmt"

func testSPClient(g EdgeWeightedDigraph, s int) {
	sp := NewSP(g, s)
	for v := 0; v < g.V(); v++ {
		fmt.Printf("%d to %d (%.2f): ", s, v, sp.DistTo(v))
		for _, e := range sp.PathTo(v) {
			fmt.Print(e, " ")
		}
		fmt.Println()
	}
}
