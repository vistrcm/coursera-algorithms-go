package union_find

//QuickFindUF Quick-Find implementation
type QuickFindUF struct {
	id []int
}

//New initializes a new QuickFindUF instance
func New(N int) QuickFindUF {
	id := make([]int, N)
	for pos, _ := range id {
		id[pos] = pos
	}

	return QuickFindUF{id: id}
}

//Connected checks if p and q are connected.
func (uf *QuickFindUF) Connected(p, q int) bool {
	return uf.id[p] == uf.id[q]
}


//Union connects p and q
func (uf *QuickFindUF) Union(p, q int) {
	pid := uf.id[p]
	qid := uf.id[q]
	for pos, elem := range uf.id {
		if elem == pid {
			uf.id[pos] = qid
		}
	}
}
