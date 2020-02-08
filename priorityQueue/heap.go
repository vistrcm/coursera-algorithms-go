package priorityQueue


//HeapPQ simple implementation of a Priority queue
type HeapPQ struct {
	items []*Item
	N     int
}

func NewHeapPQ(capacity int) *HeapPQ {
	items := make([]*Item, capacity+1)
	return &HeapPQ{items: items}
}

//IsEmpty is true if no more elements in the UnorderedMaxPQ
func (pq HeapPQ) IsEmpty() bool {
	return pq.N == 0
}

func (pq *HeapPQ)Insert(x Item) {
	pq.N++
	pq.items[pq.N] = &x
	pq.swim(pq.N)
}

func (pq HeapPQ) exch(a, b int) {
	pq.items[a], pq.items[b] = pq.items[b], pq.items[a]
}

func (pq HeapPQ) less(a, b int) bool {
	return pq.items[a].Value < pq.items[b].Value
}

func (pq *HeapPQ)DelMax() Item {
	max := pq.items[1]
	pq.exch(1, pq.N)
	pq.N--
	pq.sink(1)
	pq.items[pq.N+1] = nil // prevent loitering
	return *max
}

func (pq *HeapPQ)swim(k int){
	for k > 1 && pq.less(k/2, k) {
		pq.exch(k, k/2)
		k = k/2
	}

}

func (pq *HeapPQ)sink(k int){
	for 2*k <=pq.N {
		j := 2*k
		if j < pq.N && pq.less(j, j+1) {j++}
		if !pq.less(k, j) {break}
		pq.exch(k, j)
		k = j
	}
}

