package graph

import (
	"container/heap"
	"math"
)

// An Item is something we manage in a priority queue.
type item struct {
	vertex int
	dist   float64 // The value of the item; arbitrary.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *item, dist float64) {
	item.dist = dist
	heap.Fix(pq, item.index)
}

type DijkstraSP struct {
	edgeTo []*DirectedEdge
	distTo []float64
	pq     PriorityQueue
	inPQ   map[int]*item
}

func (dst *DijkstraSP) relax(e DirectedEdge) {
	v := e.From()
	w := e.To()
	if dst.distTo[w] > dst.distTo[v]+e.Weight() {
		dst.distTo[w] = dst.distTo[v] + e.Weight()
		dst.edgeTo[w] = &e
		if dst.pqContains(w) {
			dst.pqDecreaseKey(w, dst.distTo[w])
		} else {
			dst.heapPush(&item{
				vertex: w,
				dist:   dst.distTo[w],
			})
		}
	}
}

func (dst *DijkstraSP) heapPush(i *item) {
	dst.inPQ[i.vertex] = i
	heap.Push(&dst.pq, i)
}

func (dst *DijkstraSP) pqContains(w int) bool {
	item := dst.inPQ[w]
	return item != nil
}

func (dst *DijkstraSP) heapPop() *item {
	i := heap.Pop(&dst.pq).(*item)
	dst.inPQ[i.vertex] = nil
	return i
}

func (dst *DijkstraSP) pqDecreaseKey(w int, dist float64) {
	item := dst.inPQ[w]
	dst.pq.update(item, dist)
}

func (dst *DijkstraSP) DistTo(v int) float64 {
	return dst.distTo[v]
}

func (dst *DijkstraSP) PathTo(v int) []*DirectedEdge {
	var stack []*DirectedEdge
	var path []*DirectedEdge
	for e := dst.edgeTo[v]; e != nil; e = dst.edgeTo[e.From()] {
		stack = append(stack, e)
	}

	// now reverse stack to get path
	for len(stack) > 0 {
		n := len(stack) - 1 // Top element
		path = append(path, stack[n])
		stack = stack[:n] // Pop
	}
	return path
}

func NewDijkstraSP(g EdgeWeightedDigraph, s int) DijkstraSP {
	dsp := DijkstraSP{
		edgeTo: make([]*DirectedEdge, g.V()),
		distTo: make([]float64, g.V()),
		pq:     make(PriorityQueue, 0),
		inPQ:   make(map[int]*item),
	}
	heap.Init(&dsp.pq)

	for i := range dsp.distTo {
		dsp.distTo[i] = math.Inf(1)
	}
	dsp.distTo[s] = 0.0

	dsp.heapPush(&item{
		vertex: s,
		dist:   0,
	})

	// Take the items out; they arrive in decreasing priority order.
	for dsp.pq.Len() > 0 {
		item := dsp.heapPop()
		for _, e := range g.Adj(item.vertex) {
			dsp.relax(e)
		}
	}

	return dsp
}
