package graph

//DigraphInterface represents interface for directed graphs
type DigraphInterface interface {
	Interface
	reverse() DigraphInterface
}

//Digraph directed graph implementation
type Digraph struct {
	v   int     // amount of vertices
	e   int     // amount of edges
	adj [][]int // adjacency list
}

//NewDigraph create digraph with v vertices
func NewDigraph(v int) Graph {
	adj := make([][]int, v)
	return Graph{
		v:   v,
		adj: adj,
	}
}

//AddEdge add directed edge from v to w
func (g *Digraph) AddEdge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.e++
}

//Adj return edges adjusted to v
func (g *Digraph) Adj(v int) []int {
	return g.adj[v]
}

//V return amount of vertices
func (g *Digraph) V() int {
	return g.v
}

//E return amount of edges
func (g *Digraph) E() int {
	return g.e
}
