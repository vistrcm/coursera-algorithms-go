package graph

import "fmt"

type DirectedEdge struct {
	v      int
	w      int
	weight float64
}

func NewDirectedEdge(v, w int, weight float64) DirectedEdge {
	return DirectedEdge{
		v:      v,
		w:      w,
		weight: weight,
	}
}

func (e DirectedEdge) From() int {
	return e.v
}

func (e DirectedEdge) To() int {
	return e.w
}

func (e DirectedEdge) Weight() float64 {
	return e.weight
}

func (e DirectedEdge) String() string {
	return fmt.Sprintf("%d->%d(%.2f)", e.v, e.w, e.weight)
}

type EdgeWeightedDigraph struct {
	v   int
	adj [][]DirectedEdge
}

func NewEdgeWeightedDigraph(v int) EdgeWeightedDigraph {
	var adj = make([][]DirectedEdge, v)
	for i, _ := range adj {
		adj[i] = make([]DirectedEdge, 0)
	}

	return EdgeWeightedDigraph{
		v:   v,
		adj: adj,
	}
}

func (g *EdgeWeightedDigraph) AddEdge(e DirectedEdge) {
	v := e.From()
	g.adj[v] = append(g.adj[v], e)
}

func (g *EdgeWeightedDigraph) Adj(v int) []DirectedEdge {
	aux := make([]DirectedEdge, len(g.adj[v]))
	copy(aux, g.adj[v])
	return aux
}

//V return amount of vertices
func (g *EdgeWeightedDigraph) V() int {
	return g.v
}
