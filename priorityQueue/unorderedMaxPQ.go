package priorityQueue

type Comparable interface {
	compare(other Comparable) int
}

type Item struct {
	Value string
}

//UnorderedMaxPQ simple implementation of a Priority queue
type UnorderedMaxPQ []*Item

//NewUnorderedMaxPQ create new instance of UnorderedMaxPQ
func NewUnorderedMaxPQ(capacity int) *UnorderedMaxPQ{
	pq := make(UnorderedMaxPQ, 0, capacity)
	return &pq
}

//IsEmpty is true if no more elements in the UnorderedMaxPQ
func (pq UnorderedMaxPQ)IsEmpty() bool {
	return len(pq) == 0
}

//Insert key to the UnorderedMaxPQ
func (pq *UnorderedMaxPQ)Insert(key Item) {
	*pq = append(*pq, &key)
}

func (pq UnorderedMaxPQ)less(a, b int) bool{
	return pq[a].Value < pq[b].Value
}

func (pq UnorderedMaxPQ)swap(a, b int){
	pq[a], pq[b] = pq[b], pq[a]
}

//DelMax delete's maximum Item
func (pq *UnorderedMaxPQ)DelMax() Item {
	max := 0
	for i := 1; i< len(*pq); i++ {
		if pq.less(max, i) {max = i}
	}
	pq.swap(max, len(*pq) - 1)
	old := *pq
	item := old[len(old)-1]
	*pq = old[0: len(old)-1]
	return *item
}