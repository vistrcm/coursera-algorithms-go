package graph

type Interface interface {
	AddEdge(v, w int)
	Adj(v int) []int
}

type Graph struct {
	v   int     // amount of vertices
	adj [][]int // adjacency list
}

func NewGraph(v int) Graph {
	adj := make([][]int, v)
	return Graph{
		v:   v,
		adj: adj,
	}
}

func (g *Graph) AddEdge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
}

func (g *Graph) Adj(v int) []int {
	return g.adj[v]
}
