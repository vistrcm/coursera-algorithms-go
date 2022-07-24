package graph

import "fmt"

func ExampleNewDijkstraSP() {
	s := 0
	g := buildGraph()
	sp := NewDijkstraSP(g, s)

	for v := 0; v < g.V(); v++ {
		if len(sp.PathTo(v)) == 0 {
			continue
		}
		fmt.Printf("%d to %d (%.2f): ", s, v, sp.DistTo(v))
		for _, e := range sp.PathTo(v) {
			fmt.Print(e, ".")
		}
		fmt.Println()
	}

	//Output:
	// 0 to 1 (5.00): 0->1(5.00).
	// 0 to 2 (14.00): 0->4(9.00).4->5(4.00).5->2(1.00).
	// 0 to 3 (17.00): 0->4(9.00).4->5(4.00).5->2(1.00).2->3(3.00).
	// 0 to 4 (9.00): 0->4(9.00).
	// 0 to 5 (13.00): 0->4(9.00).4->5(4.00).
	// 0 to 6 (25.00): 0->4(9.00).4->5(4.00).5->2(1.00).2->6(11.00).
	// 0 to 7 (8.00): 0->7(8.00).
}

func buildGraph() EdgeWeightedDigraph {
	g := NewEdgeWeightedDigraph(8)
	g.AddEdge(DirectedEdge{v: 0, w: 1, weight: 5.0})
	g.AddEdge(DirectedEdge{v: 0, w: 4, weight: 9.0})
	g.AddEdge(DirectedEdge{v: 0, w: 7, weight: 8.0})
	g.AddEdge(DirectedEdge{v: 1, w: 2, weight: 12.0})
	g.AddEdge(DirectedEdge{v: 1, w: 3, weight: 15.0})
	g.AddEdge(DirectedEdge{v: 1, w: 7, weight: 4.0})
	g.AddEdge(DirectedEdge{v: 2, w: 3, weight: 3.0})
	g.AddEdge(DirectedEdge{v: 2, w: 6, weight: 11.0})
	g.AddEdge(DirectedEdge{v: 3, w: 6, weight: 9.0})
	g.AddEdge(DirectedEdge{v: 4, w: 5, weight: 4.0})
	g.AddEdge(DirectedEdge{v: 4, w: 6, weight: 20.0})
	g.AddEdge(DirectedEdge{v: 4, w: 7, weight: 5.0})
	g.AddEdge(DirectedEdge{v: 5, w: 2, weight: 1.0})
	g.AddEdge(DirectedEdge{v: 5, w: 6, weight: 13.0})
	g.AddEdge(DirectedEdge{v: 7, w: 5, weight: 6.0})
	g.AddEdge(DirectedEdge{v: 7, w: 2, weight: 7.0})
	return g
}
