package graph

type SP struct {
	distTo []float64
	edgeTo []*DirectedEdge
}

func NewSP(g EdgeWeightedDigraph, s int) SP {
	// TODO: implement
	panic("not implemented")
}

func (sp *SP) DistTo(v int) float64 {
	return sp.distTo[v]
}

func (sp *SP) PathTo(v int) []*DirectedEdge {
	var path []*DirectedEdge
	for e := sp.edgeTo[v]; e != nil; e = sp.edgeTo[e.From()] {
		path = append(path, e)
	}
	return path
}

func (sp *SP) relax(e DirectedEdge) { //nolint: unused
	v := e.From()
	w := e.To()
	if sp.distTo[w] > sp.distTo[v]+e.weight {
		sp.distTo[w] = sp.distTo[v] + e.weight
		sp.edgeTo[w] = &e
	}
}
