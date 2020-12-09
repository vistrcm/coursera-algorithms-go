package graph

//DepthFirstSearch do a DepthFirstSearch in a graph
type DepthFirstSearch struct {
	marked []bool
}

func (s *DepthFirstSearch) dfs(g Interface, v int) {
	s.marked[v] = true
	for _, w := range g.Adj(v) {
		if !s.marked[w] {
			s.dfs(g, w)
		}
	}
}

//Visited is vertices v visited on this DFS
func (s *DepthFirstSearch) Visited(v int) bool {
	return s.marked[v]
}

//NewDepthFirstSearch create instance of DepthFirstSearch
func NewDepthFirstSearch(g Interface, s int) {
	instance := &DepthFirstSearch{marked: make([]bool, g.V())}
	instance.dfs(g, s)
}
