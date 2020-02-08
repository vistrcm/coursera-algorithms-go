package priorityQueue

type Comparable interface {
	compare(other Comparable) int
}

type Item struct {
	Value string
}

//UnorderedMaxPQ simple implementation of a Priority queue
type UnorderedMaxPQ struct {
	items []*Item
	N     int
}

//NewUnorderedMaxPQ create new instance of UnorderedMaxPQ
func NewUnorderedMaxPQ(capacity int) *UnorderedMaxPQ {
	items := make([]*Item, capacity)
	return &UnorderedMaxPQ{items: items}
}

//IsEmpty is true if no more elements in the UnorderedMaxPQ
func (pq UnorderedMaxPQ) IsEmpty() bool {
	return pq.N == 0
}

//Insert key to the UnorderedMaxPQ
func (pq *UnorderedMaxPQ) Insert(key Item) {
	pq.items[pq.N] = &key
	pq.N++

}

func (pq UnorderedMaxPQ) less(a, b int) bool {
	return pq.items[a].Value < pq.items[b].Value
}

func (pq UnorderedMaxPQ) swap(a, b int) {
	pq.items[a], pq.items[b] = pq.items[b], pq.items[a]
}

//DelMax delete's maximum Item
func (pq *UnorderedMaxPQ) DelMax() Item {
	max := 0
	for i := 1; i < pq.N; i++ {
		if pq.less(max, i) {
			max = i
		}
	}
	pq.swap(max, pq.N-1)

	pq.N--
	item := pq.items[pq.N]
	return *item
}
