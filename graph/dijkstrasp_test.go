package graph

import "fmt"

func ExampleNewDijkstraSP() {
	g := buildGraph()
	dsp := NewDijkstraSP(g, 0)

	for _, r := range dsp.PathTo(6){
		fmt.Printf("%d -> %d.", r.From(), r.To())
	}
	fmt.Println()

	for _, r := range dsp.PathTo(3){
		fmt.Printf("%d -> %d.", r.From(), r.To())
	}
	fmt.Println()

	for _, r := range dsp.PathTo(7){
		fmt.Printf("%d -> %d.", r.From(), r.To())
	}

	fmt.Println()

	//Output:
	// 0 -> 4.4 -> 5.5 -> 2.2 -> 6.
	// 0 -> 4.4 -> 5.5 -> 2.2 -> 3.
	// 0 -> 7.
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
