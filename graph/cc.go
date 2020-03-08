package graph

//CC data structure to count connected components
type CC struct {
	marked []bool
	id     []int
	count  int
}

func NewCC(g Interface) *CC {
	marked := make([]bool, g.V())
	cc := &CC{
		marked: marked,
		id:     make([]int, g.V()),
		count:  0,
	}

	for v := 0; v < g.V(); v++ {
		if !cc.marked[v] {
			cc.dfs(g, v)
			cc.count++
		}
	}

	return cc
}

//Count return amount of connected components
func (cc CC) Count() int {
	return cc.count
}

//Id of component containing v
func (cc CC) Id(v int) int {
	return cc.id[v]
}

func (cc *CC) dfs(g Interface, v int) {
	cc.marked[v] = true
	cc.id[v] = cc.count
	for _, w := range g.Adj(v) {
		if !cc.marked[w] {
			cc.dfs(g, w)
		}
	}
}
