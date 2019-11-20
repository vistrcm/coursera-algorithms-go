package union_find

//QuickFindUF Quick-Find implementation
type QuickFindUF struct {
	id []int
}

//NewQuickFind initializes a new QuickFindUF instance
func NewQuickFind(N int) QuickFindUF {
	id := make([]int, N)
	for pos := range id {
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

//QuickUnionUF implements Quick Union algorithm
type QuickUnionUF struct {
	id []int
}

//NewQuickUnionUF initializes a new QuickUnionUF instance
func NewQuickUnionUF(N int) QuickUnionUF {
	id := make([]int, N)
	for pos, _ := range id {
		id[pos] = pos
	}

	return QuickUnionUF{id: id}
}

//root finds the root of element
func (uf *QuickUnionUF) root(i int) int {
	for i != uf.id[i] {
		i = uf.id[i]
	}
	return i
}

//Connected checks if p and q are connected.
func (uf *QuickUnionUF) Connected(p, q int) bool {
	return uf.root(p) == uf.root(q)
}

//Union connects p and q
func (uf *QuickUnionUF) Union(p, q int) {
	pRoot := uf.root(p)
	qRoot := uf.root(q)
	uf.id[pRoot] = qRoot
}
